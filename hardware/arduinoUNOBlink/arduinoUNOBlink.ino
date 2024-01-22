void setup() {
  Serial.begin(9600);
  delay(1000);

  Serial.println("## 1764Snippets");
  Serial.println("## Hardware Arduino UNO Blink");
  pinMode(13, OUTPUT);
}

void loop() {
  digitalWrite(13, HIGH);   // Turn the LED on
  delay(50);                // Keep it on for 500 milliseconds
  digitalWrite(13, LOW);    // Turn the LED off
  delay(50);                // Keep it off for 250 milliseconds
}
