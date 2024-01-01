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

	gns.AddRecord("example", "127.0.0.1", "A", 3600)

	record, ok := gns.GetRecord("example")
	if ok {
		fmt.Printf("Record: %v\n", record)
	} else {
		fmt.Println("Record not found")
	}
}
