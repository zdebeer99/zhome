package hardware

import (
	"fmt"
	"github.com/zdebeer99/zhome/pkg/hardware/firmata"
	"github.com/zdebeer99/zhome/pkg/hardware/qwikswitch"
	"github.com/zdebeer99/zhome/pkg/hardware/zioboard"
	"github.com/zdebeer99/zhome/pkg/stateengine"
)

func NewDeviceComm(deviceType string, id string, connection string) stateengine.DeviceComm {
	switch deviceType {
	case "qwikswitch":
		return qwikswitch.New(id, connection)
	case "firmata":
		return firmata.New(id, connection)
	case "zioboard":
		return zioboard.New(id, connection)
	default:
		panic(fmt.Errorf("device '%s' is a invalid device.", deviceType))
	}
}
