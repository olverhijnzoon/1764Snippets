#!/bin/bash

# "\033[H\033[2J" - Clears the terminal screen.
echo -e "\033[H\033[2J\033[34m===== contextSwitch =====\e[0m"

# Display context switches
echo "Context switches before:"
grep ctxt /proc/stat

# Display context switches again
echo "Context switches again:"
grep ctxt /proc/stat

# Trigger context switches by running a command that creates multiple processes
echo "Triggering context switches"
for i in {1..1000}; do echo $i > /dev/null & done

# Display context switches after
echo "Context switches after:"
grep ctxt /proc/stat
