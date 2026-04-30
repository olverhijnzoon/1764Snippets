.PHONY: all run analyze clean

all: analyze

run:
	@go run escapeAnalysis.go

# -m flag tells the compiler to print optimization decisions
# -m -m to print even more details
# -l disable inlining globally
analyze:
	go build -gcflags="-m -l" escapeAnalysis.go

clean:
	rm -f main
