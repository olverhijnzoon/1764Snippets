# Stage 1: Build the C++ application
FROM debian:buster as builder

RUN apt-get update && apt-get install -y \
    g++ \
    libcurl4-openssl-dev

COPY . /app
WORKDIR /app

RUN g++ -Wall -lcurl -o mailer main.cpp

# Stage 2: The final Docker image
FROM debian:buster

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    msmtp \
    msmtp-mta

# Install MailHog
RUN curl -o /usr/local/bin/mailhog \
    https://github.com/mailhog/MailHog/releases/download/v1.0.1/MailHog_linux_amd64 && \
    chmod +x /usr/local/bin/mailhog

# Copy the compiled binary from the builder image
COPY --from=builder /app/mailer /usr/local/bin/mailer

# Start MailHog in the background and then start the mailer application
CMD /usr/local/bin/mailhog & /usr/local/bin/mailer
