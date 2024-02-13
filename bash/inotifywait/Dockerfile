FROM ubuntu:latest

# Install required packages
RUN apt-get update && apt-get install -y inotify-tools

# Set the working directory inside the container
WORKDIR /app

# Copy the script into the container
COPY ./inotifywait.sh .

# Make the script executable
RUN chmod +x inotifywait.sh

# Command to run the script (adjust as needed)
CMD ["./inotifywait.sh"]
