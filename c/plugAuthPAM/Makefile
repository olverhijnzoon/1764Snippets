BINARY=pluggable_authentication
CC=gcc
CFLAGS=-Wall
LDFLAGS=-lpam -lpam_misc # Add linker flags for PAM and PAM_MISC libraries

all: build_docker run_docker

build_docker:
	docker build -t $(BINARY) .

run_docker:
	docker run -it --rm --privileged $(BINARY)

$(BINARY)X: $(BINARY).c
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)

run_c: $(BINARY)X
	./$(BINARY)X

build_c: $(BINARY)X

clean_c:
	rm -f $(BINARY)X
