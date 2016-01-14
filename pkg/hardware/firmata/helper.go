package firmata

import (
	"fmt"
	"github.com/zdebeer99/go-firmata"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
)

const (
	DIGITAL        byte = 0
	ANALOG         byte = 1
	PIN_MODE_DHT22      = 0x0C
	DHT_CONFIG          = 0x01
	DHT_DATA            = 0x02
)

func pinMode(channelType string) (mode firmata.PinMode, err error) {
	switch channelType {
	case "boolout":
		mode = firmata.Output
	case "boolin":
		mode = firmata.Input
	case "analogout":
		mode = firmata.Output
	case "analogin":
		mode = firmata.Input
	default:
		err = fmt.Errorf("Invalid channel Type '%s'", channelType)
	}
	return mode, err
}

// pintType returns a value indicating if the pin is a digital - 0 or analog - 1 pin.
func pinType(chType string) (ptype byte, err error) {
	switch chType {
	case "boolout":
		ptype = DIGITAL
	case "boolin":
		ptype = DIGITAL
	case "analogout":
		ptype = ANALOG
	case "analogin":
		ptype = ANALOG
	default:
		err = fmt.Errorf("Invalid channel Type '%s'", chType)
	}
	return ptype, err
}

func encodeValue(value string) (int, error) {
	switch value {
	case "on":
		return 1, nil
	case "off":
		return 0, nil
	default:
		return se.StrToInt(value)
	}
}

func decodeValue(value string) string {
	switch value {
	case "100":
		return "on"
	case "0":
		return "off"
	default:
		return value
	}
}
