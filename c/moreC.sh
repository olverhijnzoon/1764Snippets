#!/bin/bash

# Define the snippet name
snippetname="moreC"

# Create the folder
mkdir "$snippetname"

# Go to the folder
cd "$snippetname"

# Create main.c file
cat <<EOF >$snippetname.c
#include <stdio.h>

int main() {
    
    printf("1764Snippets\n");
    printf("C $snippetname\n");
    
}
EOF

# Create Makefile
cat <<EOF >Makefile
BINARY=$snippetname

run_c: build_c
	./\$(BINARY)X

build_c:
	gcc -o \$(BINARY)X \$(BINARY).c

clean_c:
	rm -f \$(BINARY)
EOF

make
