DOCKER_CONTAINER_NAME = resource-monitor-container
DOCKER_IMAGE_NAME = resource-monitor-image
SCRIPT_NAME = resourceMonitor.sh

run_monitor: docker-build
	@echo "Makefile for Docker container running resource monitoring and quota management scripts"
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

run_shell: docker-build
	docker run -it --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) /bin/bash

docker-clean:
	docker rmi $(DOCKER_IMAGE_NAME)
