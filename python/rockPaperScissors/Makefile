BINARY=rockPaperScissors

run_py: build_py
	python3 $(BINARY).py

build_py:
	pip3 install tensorflow

clean_py:
	rm -f $(BINARY).py

clean_binaries:
	rm -f ./*pyc
