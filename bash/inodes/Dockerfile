FROM ubuntu:latest
RUN apt-get update && apt-get install -y     coreutils
WORKDIR /app
COPY ./inodes.sh .
RUN chmod +x inodes.sh
