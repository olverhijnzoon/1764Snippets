DOCKER_CONTAINER_NAME=contextSwitch-container
DOCKER_IMAGE_NAME=contextswitch
SCRIPT_NAME=contextSwitch.sh

run: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

