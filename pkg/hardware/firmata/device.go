package firmata

import (
	"fmt"
	"github.com/zdebeer99/go-firmata"
	"log"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
)

type pinInfo struct {
	address string
	chType  string
	pin     byte
	mode    firmata.PinMode
	ptype   byte
}

type FirmataDevice struct {
	id         string
	Connection string
	board      *firmata.FirmataClient
	pins       map[string]pinInfo
	status     *se.Status
}

func New(id string, Connection string) *FirmataDevice {
	dev := &FirmataDevice{
		id,
		Connection,
		nil,
		make(map[string]pinInfo),
		se.NewStatus(),
	}
	var err error
	log.Println("Initialze Firmata")
	dev.board, err = firmata.NewClient(Connection, 57600)
	if err != nil {
		dev.status.SetError(fmt.Errorf("Arduino Firmata.New() Connection '%s' Failed. Error: %s", Connection, err))
		log.Println("Arduino Firmata Board Error", err)
	} else {
		go dev.eventHandler()
		log.Println("Arduino Firmata")
	}
	log.Println("Firmata Loaded! Done")
	log.Println("-----------------------")
	log.Println("Testing DHT22")
	dev.board.SetPinMode(8, firmata.DHT22)
	dev.board.ReadDHT22(8)
	log.Println("-----------------------")
	return dev
}

func (this *FirmataDevice) eventHandler() {
	for msg := range this.board.GetSerialData() {
		log.Println("---Event", msg)
	}
}

func (this *FirmataDevice) ID() string {
	return this.id
}

func (this *FirmataDevice) OK() bool {
	return this.status.IsOk()
}

func (this *FirmataDevice) RegisterChannel(is string, address string, chType string) {
	pin, err := se.StrToByte(address)
	if err != nil {
		err = fmt.Errorf("ChannelSetup Invalid address '%s'. address must be a pin number.", err)
		log.Println(err)
		return
	}
	mode, err := pinMode(chType)
	if err != nil {
		err = fmt.Errorf("ChannelSetup Invalid Pin Mode '%s'.", err)
		log.Println(err)
		return
	}
	ptype, err := pinType(chType)
	if err != nil {
		err = fmt.Errorf("ChannelSetup Invalid Pin Type '%s'.", err)
		log.Println(err)
		return
	}
	this.pins[address] = pinInfo{address, chType, pin, mode, ptype}
	if this.OK() {
		log.Println("SetPinMode", pin, mode)
		this.board.SetPinMode(pin, mode)
	}
	return
}

func (this *FirmataDevice) Status() *se.Status {
	return this.status
}

func (this *FirmataDevice) Start() {
}

func (this *FirmataDevice) Stop() {
	this.board.Close()
}

func (this *FirmataDevice) GetValue(address string) (se.ValueMap, error) {
	return se.NewValueMap(""), nil
}

func (this *FirmataDevice) SetValue(address string, value se.ValueMap) error {
	wval, err := encodeValue(value.Value())
	if err != nil {
		err = fmt.Errorf("Firmata.SetValue() Invalid Value '%s', error: %s", value, err)
		log.Println(err)
		return err
	}
	if !this.OK() {
		return fmt.Errorf("arduino device %s is in error state. %s", address, value)
	}
	if pin, ok := this.pins[address]; ok {
		if pin.mode == firmata.Output {
			switch pin.ptype {
			case DIGITAL:
				log.Println("Write", pin.pin, wval)
				err = this.board.DigitalWrite(pin.pin, wval == 1)
				if err != nil {
					return fmt.Errorf("Firmata.SetValue() write failed. %s", err)
				}
			case ANALOG:
				//this.board.WriteAnalog(pin.pin, wval)
			default:
				return fmt.Errorf("Firmata.SetValue() pin type Invalid.")
			}
		} else {
			return fmt.Errorf("Firmata.SetValue() pin mode Invalid. Can only set a value of a output pin.")
		}
	}
	return nil
}

func (this *FirmataDevice) RegisterEventHandler(handler se.EventHandler) {

}
