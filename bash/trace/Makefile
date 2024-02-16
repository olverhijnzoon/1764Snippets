DOCKER_CONTAINER_NAME=trace-container
DOCKER_IMAGE_NAME=trace
SCRIPT_NAME=trace.sh

run: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

all:
	@echo "Makefile for Docker container running strace and ltrace scripts"

trace_ls: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) ls

trace_top: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) top

clean:
	docker stop $(DOCKER_CONTAINER_NAME)
	docker rm $(DOCKER_CONTAINER_NAME)

clean-image:
	docker rmi $(DOCKER_IMAGE_NAME)