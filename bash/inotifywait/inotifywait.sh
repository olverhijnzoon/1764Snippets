#!/bin/bash

WATCHED="/tmp/"

# Run inotifywait in the background
inotifywait -m -r -e attrib "$WATCHED" | 
while read path action file; do
    echo "File $file in directory $path was modified"
done