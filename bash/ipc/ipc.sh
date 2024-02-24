#!/bin/bash

# Creating shared memory segments
ipcmk -M 1000
ipcmk -M 1764

# Creating message queues
ipcmk -Q
ipcmk -Q 

# Creating semaphore arrays
ipcmk -S 10 # semaphore array with 10 semaphores
ipcmk -S 5

echo "Listing IPC resources..."
ipcs