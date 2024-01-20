BINARY=LLDBDebugger

run_c: build_c
	./$(BINARY)X

build_c:
	gcc -o $(BINARY)X $(BINARY).c

debug_c:
	brew install llvm
	gcc -g -o $(BINARY)X $(BINARY).c
	lldb ./$(BINARY)X
	
	# (lldb) r
	# (lldb) p ptr
	# (int *) 0x0000000000000000

clean_c:
	rm -f $(BINARY)
