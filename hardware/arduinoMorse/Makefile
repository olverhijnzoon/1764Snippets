BOARD := arduino:avr:uno
PORT := /dev/cu.usbmodem1201
INO := arduinoMorse.ino

.PHONY: compile upload clean

upload: compile
	arduino-cli upload -p $(PORT) --fqbn $(BOARD) $(INO)

compile:
	arduino-cli compile --fqbn $(BOARD) $(INO)

clean:
	rm -f $(INO).hex
	rm -f $(INO).elf

prep:
	brew install arduino-cli
	arduino-cli board list
