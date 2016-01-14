package stateengine

import ()

const (
	BOOLOUT   string = "boolout"
	BOOLIN    string = "boolin"
	ANALOGOUT string = "analogout"
	ANALOGIN  string = "analogin"
)

type EventHandler func(eventName string, id string, arguments interface{})

type DeviceComm interface {
	ID() string
	RegisterChannel(id string, address string, chType string)
	RegisterEventHandler(handler EventHandler)
	Start()
	Stop()
	Status() *Status
	GetValue(chAddress string) (ValueMap, error)
	SetValue(chAddress string, value ValueMap) error
}

type ValueMap map[string]string

func NewValueMap(value string) ValueMap {
	r1 := make(ValueMap)
	r1["value"] = value
	return r1
}

func (this ValueMap) Value() string {
	if val1, ok := this["value"]; ok {
		return val1
	}
	return ""
}

type Status struct {
	StatusCode int
	StatusText string
}

func NewStatus() *Status {
	return &Status{}
}

func (this *Status) SetError(err error) {
	if err != nil {
		this.SetErrorStr(err.Error())
	}
}

func (this *Status) SetErrorStr(err string) {
	this.StatusCode = 1
	this.StatusText = err
}

func (this *Status) SetOk(message string) {
	this.StatusCode = 0
	this.StatusText = message
}

func (this *Status) IsOk() bool {
	return this.StatusCode == 0
}

type Logger interface {
	Write(msg string)
}
