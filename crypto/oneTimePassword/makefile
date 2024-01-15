BINARY=oneTimePassword
REDIS_CONTAINER_NAME=redis_for_golangtotp

run_go: run_redis
	./$(BINARY)X
	make stop_redis

run_redis: build_go
	# Pull the Redis image if not present
	docker pull redis:latest
	# Start Redis container
	docker run --name $(REDIS_CONTAINER_NAME) -d -p 6379:6379 redis:latest

build_go:
	go build -o $(BINARY)X $(BINARY).go

stop_redis:
	# Stop Redis container
	docker stop $(REDIS_CONTAINER_NAME)
	# Remove Redis container
	docker rm $(REDIS_CONTAINER_NAME)

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X

.PHONY: build_go run_redis stop_redis run_go clean_go clean_binaries
