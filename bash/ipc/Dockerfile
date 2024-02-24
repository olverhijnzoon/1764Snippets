FROM ubuntu:latest
RUN apt-get update && apt-get install -y coreutils
WORKDIR /app
COPY ./ipc.sh .
RUN chmod +x ipc.sh
