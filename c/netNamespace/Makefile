BINARY=netnamespace
CC = gcc
CFLAGS = -Wall

# The --privileged flag is required to allow the Docker container to create new namespaces.
run:
	docker build -t $(BINARY) .
	docker run -it --rm --privileged $(BINARY)
	# docker run -it --rm --privileged $(BINARY) bash

all: $(BINARY)X

$(BINARY)X: $(BINARY).c
	${CC} ${CFLAGS} -o $@ $<

run_c: build_c
	./$(BINARY)X

build_c:
	gcc -o $(BINARY)X $(BINARY).c

clean_c:
	rm -f $(BINARY)X
