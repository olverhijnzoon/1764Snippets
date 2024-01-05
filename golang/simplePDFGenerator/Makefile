BINARY=simplePDFGenerator

run_go: build_go
	./$(BINARY)

build_go:
	go build -o $(BINARY) $(BINARY).go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X
