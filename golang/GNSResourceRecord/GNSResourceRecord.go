package main

import (
	"fmt"
	"sync"
)

type ResourceRecord struct {
	Name  string
	Value string
	Type  string
	TTL   int
}

type GNS struct {
	records map[string]ResourceRecord
	mu      sync.RWMutex
}

func NewGNS() *GNS {
	return &GNS{
		records: make(map[string]ResourceRecord),
	}
}

func (g *GNS) AddRecord(name, value, recordType string, ttl int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.records[name] = ResourceRecord{Name: name, Value: value, Type: recordType, TTL: ttl}
}

func (g *GNS) GetRecord(name string) (ResourceRecord, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	record, ok := g.records[name]
	return record, ok
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang GNSResourceRecord")

	gns := NewGNS()

	// Add a record of type "PKEY" GNS zone delegation
	gns.AddRecord("publickey", "30820122300d06092a864886f70d01010105000382010f003082010a0282010100b11e30f8b4c5ba6d7f90f79e4e38c3b2cb8b2c9f2a36c63a6b8c5b5a9292487a0f98a0d69f9e39a3726e8268e0a0b3f8b5c034b1331a3f3a6c1a3a1402065574a2f2e4e7a0a0a2b6e8a8a8a2a8a0a8a2a8a8a8a2a8a0a8a2a8a8a8a2a8a0a8a2a8a8", "PKEY", 3600)
	// Add a record of type "NICK" GNS zone nickname
	gns.AddRecord("nickname", "1764SnippetsZone", "NICK", 3600)

	record, ok := gns.GetRecord("publickey")
	if ok {
		fmt.Printf("Record: %v\n", record)
	} else {
		fmt.Println("Record not found")
	}
}
