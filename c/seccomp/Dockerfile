# Use an official Ubuntu base image
FROM ubuntu:latest

# Install necessary packages
RUN apt-get update && apt-get install -y \
    gcc \
    libseccomp-dev \
    make

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the C source file and the Makefile into the container
COPY seccomp.c .
COPY Makefile .

# Build the binary when the container is built
RUN make build_c

# Command to run when starting the container
CMD ["./seccompX"]
