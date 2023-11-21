#!/bin/bash

# Snippet prep
snippet="justAnotherSnippet"
mkdir -p "$snippet"

# Dockerfile
cat > "$snippet/Dockerfile" <<EOF
FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    coreutils
WORKDIR /app
COPY ./${snippet}.sh .
RUN chmod +x ${snippet}.sh
EOF

# shell script
cat > "$snippet/${snippet}.sh" <<EOF
#!/bin/bash

# "\033[H\033[2J" - Clears the terminal screen.
echo -e "\033[H\033[2J\033[34m===== $snippet =====\e[0m"
EOF

# Makefile
cat > "$snippet/Makefile" <<EOF
DOCKER_CONTAINER_NAME=${snippet}-container
DOCKER_IMAGE_NAME=$(echo "$snippet" | tr '[:upper:]' '[:lower:]')
SCRIPT_NAME=${snippet}.sh

run: docker-build
	docker run --rm --name \$(DOCKER_CONTAINER_NAME) \$(DOCKER_IMAGE_NAME) ./\$(SCRIPT_NAME)

docker-build:
	docker build -t \$(DOCKER_IMAGE_NAME) .

EOF

# make executable and try it
chmod +x "$snippet/${snippet}.sh"
cd "$snippet"
make

