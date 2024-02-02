FROM debian:buster

# Install postfix
RUN apt-get update && apt-get install -y postfix mutt

# Clean up the apt cache to reduce image size
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Open port 25 for SMTP
EXPOSE 25

CMD ["postfix", "start-fg"]