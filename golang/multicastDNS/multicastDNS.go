package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/net/ipv4"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Multicast DNS")

	// Define the mDNS multicast address and port
	const mdnsMulticastAddr = "224.0.0.251:5353"

	// Resolve the multicast address
	multicastUDPAddr, err := net.ResolveUDPAddr("udp", mdnsMulticastAddr)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	// Open a UDP connection
	conn, err := net.ListenMulticastUDP("udp", nil, multicastUDPAddr)
	if err != nil {
		fmt.Println("Error opening UDP connection:", err)
		return
	}
	defer conn.Close()

	// Set options on the UDP connection for multicast
	p := ipv4.NewPacketConn(conn)
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		fmt.Println("Error setting control message:", err)
		return
	}

	// Send a multicast DNS query
	query := ".local"
	// found PTR records in my local e.g.
	// _companion-link._tcp.local -> MacBook
	// _rdlink._tcp.local -> iPhone
	_, err = conn.WriteTo([]byte(query), multicastUDPAddr)
	if err != nil {
		fmt.Println("Error sending query:", err)
		return
	}

	fmt.Println("mDNS query sent for domain:", query)

	// Set a timeout for listening for responses
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
		return
	}

	// Buffer to read incoming data
	buffer := make([]byte, 4096)

	// Listen for responses
	for {
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("No more responses received, stopping listener.")
				break
			}
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Process response
		fmt.Printf("Received response from %v:\n", src)
		parseDNSPacket(buffer[:n])
	}
}

func parseDNSPacket(packet []byte) {
	if len(packet) < 12 {
		fmt.Println("Invalid DNS packet")
		return
	}

	numAnswers := binary.BigEndian.Uint16(packet[6:8])
	fmt.Printf("Number of Answers: %d\n", numAnswers)

	offset := 12 // Start after the header

	// Skip question section
	for i := 0; i < int(binary.BigEndian.Uint16(packet[4:6])); i++ {
		_, off, err := readName(packet, offset)
		if err != nil {
			fmt.Println("Error reading question section:", err)
			return
		}
		offset = off + 4 // Skip QTYPE and QCLASS
	}

	// Parsing answer section
	for i := 0; i < int(numAnswers); i++ {
		name, off, err := readName(packet, offset)
		if err != nil {
			fmt.Println("Error reading answer name:", err)
			return
		}
		offset = off

		if offset+10 > len(packet) {
			fmt.Println("Invalid record header")
			return
		}

		// Process record based on type
		typ := binary.BigEndian.Uint16(packet[offset : offset+2])
		dataLen := int(binary.BigEndian.Uint16(packet[offset+8 : offset+10]))
		offset += 10

		if offset+dataLen > len(packet) {
			fmt.Println("Invalid record data length")
			return
		}

		switch typ {
		case 1: // A record
			ip := net.IPv4(packet[offset], packet[offset+1], packet[offset+2], packet[offset+3])
			fmt.Printf("A Record: %s -> %s\n", name, ip.String())
		case 28: // AAAA record
			ip := net.IP(packet[offset : offset+dataLen])
			fmt.Printf("AAAA Record: %s -> %s\n", name, ip.String())
		case 12: // PTR record
			ptrName, _, err := readName(packet, offset)
			if err != nil {
				fmt.Println("Error reading PTR record:", err)
				continue
			}
			fmt.Printf("PTR Record: %s -> %s\n", name, ptrName)
		}

		offset += dataLen
	}
}

func readName(packet []byte, offset int) (string, int, error) {
	if offset >= len(packet) {
		return "", 0, fmt.Errorf("offset out of range")
	}

	var name []string
	for {
		if offset >= len(packet) {
			return "", 0, fmt.Errorf("offset out of range")
		}

		length := int(packet[offset])
		offset++

		// Check for name compression
		if length >= 192 { // 0xC0 indicates name compression
			if offset >= len(packet) {
				return "", 0, fmt.Errorf("offset out of range")
			}
			pointer := int(packet[offset]) + (length-192)*256
			offset++
			compPart, _, err := readName(packet, pointer)
			if err != nil {
				return "", 0, fmt.Errorf("error reading compressed name: %w", err)
			}
			name = append(name, compPart)
			break
		} else if length == 0 {
			break
		} else {
			if offset+length > len(packet) {
				return "", 0, fmt.Errorf("length out of range")
			}
			name = append(name, string(packet[offset:offset+length]))
			offset += length
		}
	}
	return strings.Join(name, "."), offset, nil
}
