DOCKER_CONTAINER_NAME=inotifywait-container
DOCKER_IMAGE_NAME=inotifywait
SCRIPT_NAME=inotifywait.sh

run: docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) &
	sleep 1
	docker exec $(DOCKER_CONTAINER_NAME) touch /tmp/file1.txt
	sleep 7
	docker exec $(DOCKER_CONTAINER_NAME) touch /tmp/file7.txt
	sleep 6
	docker exec $(DOCKER_CONTAINER_NAME) touch /tmp/file6.txt
	sleep 4
	docker exec $(DOCKER_CONTAINER_NAME) touch /tmp/file4.txt

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-delete:
	docker rm -f $(DOCKER_CONTAINER_NAME)
	docker rmi $(DOCKER_IMAGE_NAME)