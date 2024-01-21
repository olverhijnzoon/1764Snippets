BINARY=threadRobotsGCD

run_c: build_c
	./$(BINARY)X

build_c:
	gcc -o $(BINARY)X $(BINARY).c

clean_c:
	rm -f $(BINARY)

debug_c:
	brew install llvm
	gcc -g -o $(BINARY)X $(BINARY).c
	lldb ./$(BINARY)X
	
	# (lldb) r
	# (lldb) thread list