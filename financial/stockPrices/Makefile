BINARY=stockPrices

run_py:
	python3 $(BINARY).py

build_py:
	pip3 install yfinance
	echo "Python script ready to be run"

clean_py:
	rm -f $(BINARY).py

clean_binaries:
	rm -f ./*pyc
