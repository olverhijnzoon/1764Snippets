#!/bin/bash

# Directory name as a variable
folder="resourceMonitor"

# The output file name
output_file="tmp"

# Check if the output file already exists and remove it
if [ -f "$output_file" ]; then
    rm "$output_file"
fi

# Iterate over each file in the directory
for file in "$folder"/*
do
    # Check if it's a regular file and not a directory
    if [ -f "$file" ]; then
        # Write the filename to the output file
        echo "### $(basename "$file")" >> "$output_file"
        # Append the content of the file to the output file
        cat "$file" >> "$output_file"
        # Optionally add a newline for separation between files
        echo -e "\n" >> "$output_file"
    fi
done

echo "All files in '$folder' have been combined into '$output_file'"
