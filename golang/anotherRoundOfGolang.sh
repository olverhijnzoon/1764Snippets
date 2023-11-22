#!/bin/bash

# Define the snippet name
snippetname="justAnotherSnippet"

# Create the folder
mkdir "$snippetname"

# Go to the folder
cd "$snippetname"

# Create main.go file
cat <<EOF >$snippetname.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang $snippetname")
}
EOF

# Create go.mod file
cat <<EOF >go.mod
module github.com/olverhijnzoon/1764Snippets/$snippetname

go 1.21
EOF

# Create Makefile
cat <<EOF >Makefile
BINARY=$snippetname

run_go: build_go
	./\$(BINARY)X

build_go:
	go build -o \$(BINARY)X \$(BINARY).go

clean_go:
	rm -f \$(BINARY)

clean_binaries:
	rm -f ./*X
EOF

make