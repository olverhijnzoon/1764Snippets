DOCKER_CONTAINER_NAME=justAnotherSnippet-container
DOCKER_IMAGE_NAME=justanothersnippet
SCRIPT_NAME=justAnotherSnippet.sh

run: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

