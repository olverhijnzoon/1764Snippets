#!/bin/bash

# Function to check if command exists
function command_exists() {
  type "$1" &> /dev/null ;
}

# Check if strace and ltrace commands exist
if ! command_exists strace || ! command_exists ltrace; then
  echo "strace and/or ltrace not found. Please install them and try again."
  exit 1
fi

# Check if a command to trace was supplied
if [ -z "$1" ]; then
  echo "Usage: $0 <command>"
  exit 1
fi

command_to_trace="$1"

# Print debug information
echo "Debug: command to trace is '$command_to_trace'"

# Run strace on the command, capturing the output
strace_output=$(strace $command_to_trace 2>&1)

# Define a function for printing with a delay and in green
function slow_print() {
  while IFS= read -r line; do
    echo -e "\e[32m$line\e[0m"
    sleep 1.764
  done <<< "$1"
}

# "\033[H\033[2J" - Clears the terminal screen.
echo -e "\033[H\033[2J\033[32m===== System Calls by $command_to_trace =====\e[0m"
slow_print "$strace_output"

# Run ltrace on the command, capturing the output
ltrace_output=$(ltrace $command_to_trace 2>&1)

echo -e "\e[32m===== Library Calls by $command_to_trace =====\e[0m"
slow_print "$ltrace_output"