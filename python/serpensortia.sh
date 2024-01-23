#!/bin/bash

# Define the snippet name
snippetname="sysPath"

# Create the folder
mkdir "$snippetname"

# Go to the folder
cd "$snippetname"

# Create main.py file
cat <<EOF >$snippetname.py
import sys
print(sys.path)
EOF

# Create Makefile
cat <<EOF >Makefile
BINARY=$snippetname

run_py: build_py
	python3 \$(BINARY).py

build_py:
	# pip3 install tensorflow
	# echo "Python script ready to be run"

clean_py:
	rm -f \$(BINARY).py

clean_binaries:
	rm -f ./*pyc
EOF

make
