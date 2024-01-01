#!/bin/bash

# "\033[H\033[2J" - Clears the terminal screen.
echo -e "\033[H\033[2J\033[34m===== FTPServer =====\e[0m"

# Start the FTP server in the foreground
vsftpd
