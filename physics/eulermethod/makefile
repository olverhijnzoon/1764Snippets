BINARY=eulermethod
build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

god: build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X