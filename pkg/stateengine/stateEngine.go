package stateengine

import (
	"container/list"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

type StateEngine struct {
	running     bool
	deviceComm  map[string]DeviceComm
	channels    map[string]*Channel
	triggers    []Trigger
	schedules   []Schedule
	taskQueue   *list.List
	tickerTasks *time.Ticker
	Scheduler   *cron.Cron
}

func New() *StateEngine {
	se := &StateEngine{}
	se.deviceComm = make(map[string]DeviceComm)
	se.channels = make(map[string]*Channel)
	se.triggers = make([]Trigger, 0)
	se.schedules = make([]Schedule, 0)
	return se
}

func (this *StateEngine) Start() {
	this.running = true
	for _, dev := range this.deviceComm {
		dev.Start()
	}
	this.startScheduler()
}

func (this *StateEngine) Stop() {
	this.Scheduler.Stop()
	this.running = false
	for _, dev := range this.deviceComm {
		dev.Stop()
	}
}

func (this *StateEngine) Print() {
	fmt.Println("Devices")
	for k, _ := range this.deviceComm {
		fmt.Printf("id %s \n", k)
	}
	fmt.Println("Channels")
	for k, _ := range this.channels {
		fmt.Printf("id %s \n", k)
	}
	fmt.Println("Triggers")
	for k, val := range this.triggers {
		fmt.Println("id ", k, val)
	}
}

func (this *StateEngine) RegisterDevice(deviceId string, device DeviceComm) {
	device.RegisterEventHandler(this.eventHandler)
	this.deviceComm[deviceId] = device
}

func (this *StateEngine) RegisterChannel(channelId string, deviceId string, address string, chType string) {
	if device, ok := this.deviceComm[deviceId]; ok {
		this.channels[channelId] = newChannel(device, address, chType)
		device.RegisterChannel(channelId, address, chType)
	} else {
		log.Printf("Register Channel: Device '%s' not found.", deviceId)
	}
}

func (this *StateEngine) AddTrigger(trigger Trigger) {
	this.triggers = append(this.triggers, trigger)
}

func (this *StateEngine) AddSchedule(name, when string, commands []string) {
	this.schedules = append(this.schedules, Schedule{name, when, commands})
}

func (this *StateEngine) SetValue(channelId string, value string) ChannelValue {
	log.Println("SetValue ", channelId, value)
	if channel, ok := this.channels[channelId]; ok {
		channel.SetValue(NewValueMap(value))
		err := channel.Device.SetValue(channel.Address, NewValueMap(value))
		channel.SetError(err)
		return channel.Value
	}
	panic(fmt.Errorf("channel '%s' not found.", channelId))
}

func (this *StateEngine) RequestValue(channelId string) ChannelValue {
	if ch, ok := this.channels[channelId]; ok {
		value, err := ch.Device.GetValue(ch.Address)
		ch.SetValue(value)
		ch.SetError(err)
		log.Println("RequestValue Result", ch.Value)
		return ch.Value
	} else {
		return ChannelValue{StatusCode: 1, StatusText: fmt.Sprintf("channel '%s' not found.", channelId)}
	}
}

func (this *StateEngine) GetValue(channelId string) ChannelValue {
	if ch, ok := this.channels[channelId]; ok {
		return ch.Value
	} else {
		return ChannelValue{StatusCode: 1, StatusText: fmt.Sprintf("channel '%s' not found.", channelId)}
	}
}
