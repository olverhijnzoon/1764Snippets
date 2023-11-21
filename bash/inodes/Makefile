DOCKER_CONTAINER_NAME=inodes-container
DOCKER_IMAGE_NAME=inodes
SCRIPT_NAME=inodes.sh

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# target for showing inode information of '/bin' directory
inode_bin: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) /bin

# target for showing inode information of '/etc' directory
inode_etc: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) /etc