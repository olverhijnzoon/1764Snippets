DOCKER_CONTAINER_NAME=FTPServer-container
DOCKER_IMAGE_NAME=ftpserver
SCRIPT_NAME=FTPServer.sh

run: docker-build
	docker run --rm -p 21:21 --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

lftp: 
	lftp ftp://localhost
