FROM ubuntu:latest

RUN apt-get update && apt-get install -y strace ltrace

WORKDIR /app

COPY ./trace.sh .
RUN chmod +x trace.sh