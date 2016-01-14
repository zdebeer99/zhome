/*
a very basic serial protocal to set and get pin values.
the protocol works by sending a byte array, where the first byte is a command
bytes following the command is command specific.

the return packet should always be a string with the first character indicating the result type;
'a' - ack
'e' - error
'_' - information

The last character of the result must be a simicolon ';'
*/

//load dht library for communicating with the dht22 temperature sensor.
#include <dht.h>

dht DHT;

//Commands
#define CMD_PING          10
#define CMD_PIN_MODE      11
#define CMD_WRITE_DIGITAL 12
#define CMD_WRITE_ANALOG  13
#define CMD_READ_DIGITAL  14
#define CMD_READ_ANALOG   15
#define CMD_READ_DHT22    16

#define PIN_LOW  0
#define PIN_HIGH 1

#define MODE_INPUT 0
#define MODE_OUTPUT 1

void setup() {
  Serial.begin(57600);
  Serial.print("_Device ZIOBoard;");
  Serial.print("_Version 0.01;");
}

void loop() {
  delay(200);
  if (Serial.available() > 0) {
    switch (Serial.read()) {
      case CMD_PING:
        fsPing();
      case CMD_PIN_MODE:
        fsPinMode();
        break;
      case CMD_WRITE_DIGITAL:
        fsSetDigital();
        break;
      case CMD_READ_DHT22:
        fsReadDHT22();
        break;
    }
  }
}

void fsPing() {
  Serial.print("a;");
}

void fsPinMode() {
  int pin = Serial.read();
  if (pin == -1) {
    Serial.print("e;");
    return;
  }
  int smode = Serial.read();
  if (smode == -1) {
    Serial.print("e;");
    return;
  }
  switch (smode) {
    case MODE_INPUT:
      pinMode(pin, INPUT);
      Serial.print("a;");
      break;
    case MODE_OUTPUT:
      pinMode(pin, OUTPUT);
      Serial.print("a;");
      break;
    default:
      Serial.print("eInvalid pin mode;");
      break;
  }
}

void fsSetDigital() {
  int pin = Serial.read();
  if (pin == -1) {
    Serial.print("e;");
    return;
  }
  int value = Serial.read();
  if (value == -1) {
    Serial.print("e;");
    return;
  }
  digitalWrite(pin, value);
  Serial.print("a;");
}

void fsReadDHT22() {
  int pin = Serial.read();
  if (pin == -1) {
    Serial.print("e;");
    return;
  }
  int chk = DHT.read22(pin);
  switch (chk)
  {
  case DHTLIB_OK:
      break;
  case DHTLIB_ERROR_CHECKSUM:
      Serial.print("eChecksum error;");
      return;
  case DHTLIB_ERROR_TIMEOUT:
      Serial.print("eTime out error;");
      return;
  case DHTLIB_ERROR_CONNECT:
      Serial.print("eConnect error;");
      return;
  case DHTLIB_ERROR_ACK_L:
      Serial.print("eAck Low error;");
      return;
  case DHTLIB_ERROR_ACK_H:
      Serial.print("eAck High error;");
      return;
  default:
      Serial.print("eUnknown error;");
      return;
  }
  Serial.print("a");
  Serial.print(DHT.humidity, 1);
  Serial.print(",");
  Serial.print(DHT.temperature, 1);
  Serial.print(";");
}
