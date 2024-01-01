FROM ubuntu:latest
RUN apt-get update && apt-get install -y vsftpd
WORKDIR /app
COPY ./FTPServer.sh .
RUN chmod +x FTPServer.sh

# Configure vsftpd
RUN mkdir -p /var/run/vsftpd/empty
RUN echo "anonymous_enable=YES" >> /etc/vsftpd.conf