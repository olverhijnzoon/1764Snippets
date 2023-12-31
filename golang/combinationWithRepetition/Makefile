BINARY=combinationWithRepetition

run_go: build_go
	./$(BINARY)X

build_go:
	go build -o $(BINARY)X $(BINARY).go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X
