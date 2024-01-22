#!/bin/bash

# Define the snippet name
snippetname="anotherHardware"

# Create the folder
mkdir "$snippetname"

# Go to the folder
cd "$snippetname"

# Create anotherHardware.ino file
cat <<EOF >$snippetname.ino
void setup() {
  Serial.begin(9600);
  delay(1000);

  Serial.println("## 1764Snippets");
  Serial.println("## Hardware Arduino UNO");
  pinMode(13, OUTPUT);
}

void loop() {
  // Your loop code here
}
EOF

# Create Makefile
cat <<EOF >Makefile
BOARD := arduino:avr:uno
PORT := /dev/cu.usbmodem1201
INO := anotherHardware.ino

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
EOF