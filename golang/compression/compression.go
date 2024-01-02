package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Compression Standard Library")

	/*
		This code demonstrates how to compress and decompress data using gzip, and zlib. It creates a byte slice of original data, and then compresses and decompresses the data using each of the algorithms. Finally, it prints the original data and the decompressed data for each algorithm to verify that the compression and decompression worked correctly.

		This is slightly modified version of the original 42Snippets golang compression snippet "Compression Algorithms".
	*/

	// Define the input data
	input := []byte("123450afghjikop#qrtswxyz-?\n")

	// Compress the data using gzip
	var gzipBuf bytes.Buffer
	gzipWriter := gzip.NewWriter(&gzipBuf)
	if _, err := gzipWriter.Write(input); err != nil {
		fmt.Println("Error compressing data with gzip:", err)
		return
	}
	if err := gzipWriter.Close(); err != nil {
		fmt.Println("Error closing gzip writer:", err)
		return
	}

	// Compress the data using zlib
	var zlibBuf bytes.Buffer
	zlibWriter, err := zlib.NewWriterLevel(&zlibBuf, zlib.BestCompression)
	if err != nil {
		fmt.Println("Error creating zlib writer:", err)
		return
	}
	if _, err := zlibWriter.Write(input); err != nil {
		fmt.Println("Error compressing data with zlib:", err)
		return
	}
	if err := zlibWriter.Close(); err != nil {
		fmt.Println("Error closing zlib writer:", err)
		return
	}

	// Decompress the gzip-compressed data
	gzipReader, err := gzip.NewReader(&gzipBuf)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	var gzipOutput bytes.Buffer
	if _, err := io.Copy(&gzipOutput, gzipReader); err != nil {
		fmt.Println("Error decompressing data with gzip:", err)
		return
	}
	if err := gzipReader.Close(); err != nil {
		fmt.Println("Error closing gzip reader:", err)
		return
	}

	// Decompress the zlib-compressed data
	zlibReader, err := zlib.NewReader(&zlibBuf)
	if err != nil {
		fmt.Println("Error creating zlib reader:", err)
		return
	}
	var zlibOutput bytes.Buffer
	if _, err := io.Copy(&zlibOutput, zlibReader); err != nil {
		fmt.Println("Error decompressing data with zlib:", err)
		return
	}
	if err := zlibReader.Close(); err != nil {
		fmt.Println("Error closing zlib reader:", err)
		return
	}

	// Print the results
	fmt.Print("Original data:", string(input))
	fmt.Print("Gzip-decompressed data:", gzipOutput.String())
	fmt.Print("Zlib-decompressed data:", zlibOutput.String())
}
