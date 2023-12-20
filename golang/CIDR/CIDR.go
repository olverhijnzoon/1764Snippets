package main

import (
	"fmt"
	"math"
	"net"
)

// CIDRInfo holds information about each CIDR network prefix.
type CIDRInfo struct {
	NetworkPrefixLength int
	SubnetMask          string
	AvailableAddresses  int
}

func ipMaskToString(mask net.IPMask) string {
	ip := net.IP(mask).To4()
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang CIDR")

	var cidrTable []CIDRInfo

	// Generate CIDR table for prefix lengths from 0 to 32
	for prefixLength := 1; prefixLength <= 32; prefixLength++ {
		cidrInfo := CIDRInfo{
			NetworkPrefixLength: prefixLength,
			SubnetMask:          ipMaskToString(net.CIDRMask(prefixLength, 32)),
			AvailableAddresses:  int(math.Pow(2, float64(32-prefixLength))),
		}
		cidrTable = append(cidrTable, cidrInfo)
	}

	fmt.Println("CIDR\tSubnet mask\t\tAvailable addresses")
	for _, cidr := range cidrTable {
		fmt.Printf("/%-2d\t%s\t\t%v\n",
			cidr.NetworkPrefixLength, cidr.SubnetMask, cidr.AvailableAddresses)
	}
}
