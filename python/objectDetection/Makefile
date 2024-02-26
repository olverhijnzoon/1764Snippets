BINARY=objectDetectionVideo

run_py: build_py
	python3 $(BINARY).py

build_py:
	pip3 install opencv-python --break-system-packages
	pip3 install pyautogui --break-system-packages

clean_py:
	rm -f $(BINARY).py

clean_binaries:
	rm -f ./*pyc
