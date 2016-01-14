package stateengine

import (
	"time"
	"github.com/zdebeer99/zhome/pkg/config"
)

type ChannelState struct {
	Info  config.Channel `json:"info"`
	State ChannelValue   `json:"state"`
}

type Channel struct {
	Device      DeviceComm
	Address     string
	ChannelType string
	Value       ChannelValue
}

func newChannel(device DeviceComm, address string, channelType string) *Channel {
	return &Channel{
		Device:      device,
		Address:     address,
		ChannelType: channelType,
	}
}

type ChannelValue struct {
	LastUpdate time.Time `json:"lastUpdate"`
	Data       ValueMap  `json:"value"`
	StatusCode int       `json:"statusCode"`
	StatusText string    `json:"statusText"`
}

func (this *Channel) SetValue(value ValueMap) {
	this.Value.LastUpdate = time.Now()
	this.Value.Data = value
}

func (this *Channel) SetError(err error) {
	if err != nil {
		this.SetErrorStr(err.Error())
	} else {
		this.SetOk()
	}
}

func (this *Channel) SetErrorStr(err string) {
	this.Value.StatusCode = 1
	this.Value.StatusText = err
}

func (this *Channel) SetOk() {
	this.SetOkStr("")
}

func (this *Channel) SetOkStr(message string) {
	this.Value.StatusCode = 0
	this.Value.StatusText = message
}

// Trigger does something if certian conditions is met.
type Trigger struct {
	Name      string   `yaml:"Name" json:"name"`
	EventName string   `yaml:"EventName" json:"eventName"`
	Command   []string `yaml:"Command" json:"command"`
}

type Schedule struct {
	Name     string   `yaml:"Name" json:"name"`
	When     string   `yaml:"When" json:"when"`
	Commands []string `yaml:"Command" json:"command"`
}

type task struct {
	createdTime time.Time
	targetTime  time.Time
	commands    []string
}
