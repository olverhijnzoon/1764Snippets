# Use a base image with gcc installed
FROM ubuntu:latest

# Install build essentials and PAM development libraries
RUN apt-get update && \
    apt-get install -y gcc make libpam0g-dev && \
    rm -rf /var/lib/apt/lists/*

# Create a new user 'snippet1764' with home directory
RUN useradd -m snippet1764

# Set the password for the user to '42'
# Using whois package's mkpasswd for generating password hash
RUN echo 'snippet1764:42' | chpasswd

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the Makefile and source code into the container
COPY . .

# Compile the program using the Makefile
RUN make build_c

# Switch to user 'snippet1764'
USER snippet1764

# Command to run the compiled binary
CMD ["./pluggable_authenticationX"]
