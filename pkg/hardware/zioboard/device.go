package zioboard

import (
	"fmt"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
	"log"
	"strings"
	"time"
)

type pinInfo struct {
	id         string
	address    string
	chType     string
	pin        byte
	mode       byte
	ptype      byte
	configured bool
}

/*
  boardState int
    0 - new instance
    10 - running
    20 - stopped

*/
type ZIOBoardDevice struct {
	id         string
	Connection string
	board      *ZIOBoard
	pins       map[string]pinInfo
	status     *se.Status
	boardState int
}

func New(id string, connection string) *ZIOBoardDevice {
	dev := &ZIOBoardDevice{
		id,
		connection,
		nil,
		make(map[string]pinInfo),
		se.NewStatus(),
		0,
	}
	dev.status.SetErrorStr("initilizing device.")
	dev.board = NewZIOBoard(connection)
	return dev
}

func (this *ZIOBoardDevice) ID() string {
	return this.id
}

func (this *ZIOBoardDevice) OK() bool {
	if !this.board.IsOpen() {
		return false
	}
	return this.status.IsOk()
}

func (this *ZIOBoardDevice) StatusText() string {
	if this.OK() {
		return ""
	}
	if !this.board.IsOpen() {
		return "The comminication port is not open. " + this.board.state
	}
	return this.status.StatusText
}

func (this *ZIOBoardDevice) RegisterChannel(id string, address string, chType string) {
	pin, err := se.StrToByte(address)
	if err != nil {
		err = fmt.Errorf("ZIOBoard '%s' ChannelSetup Invalid address '%s'. address must be a pin number. %s", id, address, err)
		log.Println(err)
		return
	}
	mode, err := pinMode(chType)
	if err != nil {
		err = fmt.Errorf("ZIOBoard '%s' ChannelSetup Invalid Pin Mode '%s'. %s", id, chType, err)
		log.Println(err)
		return
	}
	ptype, err := pinType(chType)
	if err != nil {
		err = fmt.Errorf("ZIOBoard '%s' ChannelSetup Invalid Pin Type '%s'. %s", id, chType, err)
		log.Println(err)
		return
	}
	this.pins[address] = pinInfo{id, address, chType, pin, mode, ptype, false}
	return
}

func (this *ZIOBoardDevice) Status() *se.Status {
	return this.status
}

func (this *ZIOBoardDevice) Start() {
	go this.worker()
}

func (this *ZIOBoardDevice) Stop() {
	this.board.Close()
}

func (this *ZIOBoardDevice) GetValue(address string) (value se.ValueMap, err error) {
	value = se.NewValueMap("")
	// if !this.OK() {
	// 	err = fmt.Errorf("arduino device %s is in error state. %s", address, value)
	// 	return
	// }
	if pin, ok := this.pins[address]; ok {
		if pin.mode == MODE_INPUT {
			if pin.chType == "dht22" {
				value, err = this.readDHT22(pin.pin)
				return
			}
		}
	}
	return
}

func (this *ZIOBoardDevice) readDHT22(pin byte) (value se.ValueMap, err error) {
	valuestr, err := this.board.Send([]byte{CMD_READ_DHT22, pin})
	if err != nil {
		return
	}
	log.Println("readDHT22", valuestr, err)
	values := strings.Split(valuestr, ",")
	if len(values) > 1 {
		value = se.NewValueMap(values[1])
		value["temp"] = values[1]
		value["hum"] = values[0]
	} else {
		value = se.NewValueMap("novalue")
	}

	return
}

func (this *ZIOBoardDevice) SetValue(address string, value se.ValueMap) error {
	if !this.OK() {
		return fmt.Errorf("arduino device %s is in error state. %s", this.ID(), this.StatusText())
	}
	if pin, ok := this.pins[address]; ok {
		if pin.mode == MODE_OUTPUT {
			switch pin.ptype {
			case DATA_TYPE_DIGITAL:
				var dv1 bool
				switch value.Value() {
				case "on":
					dv1 = true
				case "off":
					dv1 = false
				default:
					return fmt.Errorf("Invalid Digital Value. `%s` expecting on or off", value)
				}
				log.Println("Write", pin.pin, dv1)
				err := this.board.WriteDigital(pin.pin, dv1)
				if err != nil {
					return fmt.Errorf("ZIOBoard write failed. %s", err)
				}
			case DATA_TYPE_ANALOG:
				//this.board.WriteAnalog(pin.pin, wval)
			default:
				return fmt.Errorf("ZIOBoard.SetValue() pin type Invalid.")
			}
		} else {
			return fmt.Errorf("ZIOBoard.SetValue() pin mode Invalid. Can only set a value of a output pin.")
		}
	}
	return nil
}

func (this *ZIOBoardDevice) RegisterEventHandler(handler se.EventHandler) {

}

func (this *ZIOBoardDevice) worker() {
	for this.boardState < 20 {
		switch this.boardState {
		case 0:
			this.boardState = this.initialize()
		case 1:
			this.boardState = this.initPins()
		case 2:
			this.status.SetOk("initialized.")
			this.boardState = 10
		case 10:
			this.boardState = this.healthCheck()
		}
		time.Sleep(1 * time.Second)
	}
}

func (this *ZIOBoardDevice) initialize() int {
	err := this.board.Open()
	if err != nil {
		this.status.SetError(fmt.Errorf("Arduino ZIOBoard. Connection '%s' Failed. Error: %s", this.Connection, err))
		log.Println("Arduino ZIOBoard Board Connection Error", err)
		time.Sleep(10 * time.Second)
		return 0
	}
	return 1
}

func (this *ZIOBoardDevice) initPins() int {
	if this.board.IsOpen() {
		for _, pin := range this.pins {
			if pin.ptype != DATA_TYPE_IGNORE && !pin.configured {
				log.Println("SetPinMode", pin.pin, pin.mode)
				err := this.board.PinMode(pin.pin, pin.mode)
				if err != nil {
					log.Printf("SetPinMode Failed pin %v Error %s", pin.pin, err)
					continue
				}
				pin.configured = true
			}
		}
	} else {
		return 0
	}
	return 2
}

func (this *ZIOBoardDevice) healthCheck() int {
	if !this.board.IsOpen() {
		return 0
	}
	time.Sleep(10 * time.Second)
	return 10
}
