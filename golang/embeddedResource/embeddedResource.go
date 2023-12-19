package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed embedded.png
var content embed.FS

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Embedded Resource")

	http.Handle("/", http.FileServer(http.FS(content)))
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
