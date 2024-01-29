#!/bin/bash

interval=2
count=4

echo -e "\e[H\e[2J\e[32m"
echo "--------------------------------------------"
echo "#  1764Snippets"
echo "## Bash Resource Monitor"
echo "--------------------------------------------"

# Run sar command
echo "Collecting system activity information"
sar $interval $count
echo "--------------------------------------------"

# Print CPU usage
echo "Printing CPU usage"
sar -u $interval $count
echo "--------------------------------------------"

# Print I/O statistics
echo "Printing I/O statistics"
sar -b $interval $count
echo "--------------------------------------------"

# Print memory usage statistics
echo "Printing memory usage statistics"
sar -r $interval $count
echo "--------------------------------------------"