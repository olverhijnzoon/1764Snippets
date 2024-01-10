FROM ubuntu:latest
RUN apt-get update && apt-get install -y procps time
WORKDIR /app
COPY ./contextSwitch.sh .
RUN chmod +x contextSwitch.sh
