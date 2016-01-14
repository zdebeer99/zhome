package zioboard

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strings"
	"time"
)

const (
	CMD_PIN_MODE      byte = 11
	CMD_WRITE_DIGITAL byte = 12
	CMD_WRITE_ANALOG  byte = 13
	CMD_READ_DIGITAL  byte = 14
	CMD_READ_ANALOG   byte = 15
	CMD_READ_DHT22    byte = 16

	PIN_LOW  byte = 0
	PIN_HIGH byte = 1

	MODE_INPUT  byte = 0
	MODE_OUTPUT byte = 1
)

type ZIOBoard struct {
	connection string
	port       *serial.Port
	running    bool
	reader     chan string
	state      string
}

func NewZIOBoard(connection string) *ZIOBoard {
	return &ZIOBoard{connection: connection, state: "new"}
}

func (this *ZIOBoard) Open() error {
	this.state = "opening"
	c := &serial.Config{Name: this.connection, Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		this.state = "closed"
		if s != nil {
			log.Println("Open() ZIOBoard Port Closed.")
			s.Close()
		}
		return err
	}
	//wait 2 seconds to allow for arduino to boot.
	time.Sleep(2 * time.Second)
	this.state = "open"
	this.running = true
	this.port = s
	go this.listen()
	return nil
}

func (this *ZIOBoard) IsOpen() bool {
	return this.state == "open"
}

func (this *ZIOBoard) Close() error {
	this.state = "closing"
	this.running = false
	if this.port != nil {
		err := this.port.Close()
		if err != nil {
			return err
		}
	}
	this.port = nil
	this.state = "closed"
	return nil
}

func (this *ZIOBoard) listen() {
	this.reader = make(chan string, 3)
	readbuf := make([]byte, 1)
	packet := new(bytes.Buffer)
	end := []byte(";")
	for this.running {
		_, err := this.port.Read(readbuf)
		if err != nil {
			log.Println("ZIOBoard Reead Error: ", err)
			this.Close()
			return
		}
		packet.Write(readbuf)
		if readbuf[0] == end[0] {
			this.handleIncoming(packet.String())
			packet = new(bytes.Buffer)
		}
	}
}

func (this *ZIOBoard) handleIncoming(packet string) {
	//log.Println("Packet Recieved.", packet)
	if len(packet) == 0 {
		return
	}
	if strings.HasPrefix(packet, "a") {
		this.reader <- packet
		return
	}
	if strings.HasPrefix(packet, "e") {
		this.reader <- packet
		return
	}
	log.Println("Command Recieved.", packet)
}

func (this *ZIOBoard) waitResponse() (data string, err error) {
	select {
	case cmd := <-this.reader:
		data, err = parseResponse(cmd)
	case <-time.After(1 * time.Second):
		log.Println("timeout")
		err = fmt.Errorf("Command Timeout.")
	}
	return
}

func (this *ZIOBoard) Send(data []byte) (r string, err error) {
	r = ""
	_, err = this.port.Write(data)
	if err != nil {
		return
	}
	r, err = this.waitResponse()
	return
}

func (this *ZIOBoard) PinMode(pin, mode byte) (err error) {
	_, err = this.Send([]byte{CMD_PIN_MODE, pin, mode})
	return
}

func (this *ZIOBoard) WriteDigital(pin byte, val bool) (err error) {
	var val1 byte
	if val {
		val1 = PIN_HIGH
	} else {
		val1 = PIN_LOW
	}
	_, err = this.Send([]byte{CMD_WRITE_DIGITAL, pin, val1})
	return
}

func parseResponse(packet string) (r string, err error) {
	if strings.HasPrefix(packet, "a") {
		if len(packet) > 2 {
			r = packet[1 : len(packet)-1]
		}
		return
	}
	if strings.HasPrefix(packet, "e") {
		if len(packet) > 2 {
			err = errors.New(packet[1 : len(packet)-1])
		} else {
			err = errors.New("Error recieved from arduino.")
			panic(err)
		}
		return
	}
	err = fmt.Errorf("Unkown packet recieved. %s", packet)
	return
}
