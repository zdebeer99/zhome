package zioboard

import (
	"fmt"
)

const (
	DATA_TYPE_IGNORE  byte = 0
	DATA_TYPE_DIGITAL byte = 1
	DATA_TYPE_ANALOG  byte = 2
)

func pinMode(channelType string) (mode byte, err error) {
	switch channelType {
	case "boolout":
		mode = MODE_OUTPUT
	case "boolin":
		mode = MODE_INPUT
	case "analogout":
		mode = MODE_OUTPUT
	case "analogin":
		mode = MODE_INPUT
	case "dht22":
		mode = MODE_INPUT
	default:
		err = fmt.Errorf("Invalid channel Type '%s'", channelType)
	}
	return mode, err
}

func pinType(chType string) (ptype byte, err error) {
	switch chType {
	case "boolout":
		ptype = DATA_TYPE_DIGITAL
	case "boolin":
		ptype = DATA_TYPE_DIGITAL
	case "analogout":
		ptype = DATA_TYPE_ANALOG
	case "analogin":
		ptype = DATA_TYPE_ANALOG
	case "dht22":
		ptype = DATA_TYPE_IGNORE
	default:
		err = fmt.Errorf("Invalid channel Type '%s'", chType)
	}
	return ptype, err
}
