DOCKER_CONTAINER_NAME=ipc-container
DOCKER_IMAGE_NAME=ipc
SCRIPT_NAME=ipc.sh

ipc: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)
	
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .