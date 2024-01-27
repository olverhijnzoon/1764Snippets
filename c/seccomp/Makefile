BINARY=seccomp
IMAGE_NAME=seccomp-demo

run_docker:
	docker run --rm -it $(IMAGE_NAME) /bin/bash
	# root@fad90d0293ee:/usr/src/app# ./seccompX 
	
build_docker:
	docker build -t $(IMAGE_NAME) .

clean_docker:
	docker rmi $(IMAGE_NAME)

build_c:
	gcc -o $(BINARY)X $(BINARY).c -lseccomp

run_c: build_c
	./$(BINARY)X

clean_c:
	rm -f $(BINARY)X
