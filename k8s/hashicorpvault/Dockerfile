# Start from the official Vault image
FROM vault:1.7.3

# Set the working directory
WORKDIR /vault/config

# Copy the local configuration files to the Docker image
COPY . .

# Run Vault
CMD ["vault", "server", "-config=/vault/config"]
