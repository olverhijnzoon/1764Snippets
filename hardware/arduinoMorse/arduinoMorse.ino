// Morse code representation for each character
const char* morseCode[36] = {
  ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
  "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
  "-----", ".----", "..---", "...--", "....-", ".....", "-....", "--...", "---..", "----."
};

void setup() {
  Serial.begin(9600);
  delay(1000);

  Serial.println("## 1764Snippets");
  Serial.println("## Hardware Arduino UNO Morse Code Blink");
  pinMode(13, OUTPUT);
}

void loop() {
  String message = "0505050505";
  sendMorse(message);
  delay(10000); // Wait for 3 seconds before repeating
}

void sendMorse(String message) {
  for (int i = 0; i < message.length(); i++) {
    char c = message.charAt(i);
    if (c >= 'A' && c <= 'Z') {
      transmitMorseCode(morseCode[c - 'A']);
    } else if (c >= '0' && c <= '9') {
      transmitMorseCode(morseCode[c - '0' + 26]);
    }
    delay(700); // Space between letters
  }
}

void transmitMorseCode(const char* morse) {
  for (int i = 0; morse[i] != '\0'; i++) {
    if (morse[i] == '.') {
      blinkLED(250); // Dot
    } else if (morse[i] == '-') {
      blinkLED(750); // Dash
    }
    delay(250); // Space between dots/dashes
  }
}

void blinkLED(int duration) {
  digitalWrite(13, HIGH);
  delay(duration);
  digitalWrite(13, LOW);
  delay(duration);
}
