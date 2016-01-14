package stateengine

import (
	"zdebeer/qwickswitch/server/config"
)

var State = New()
var AppConfig *config.Config

func Start() {
	State.Start()
}

func Stop() {
	State.Stop()
}

func SetValue(channelId string, value string) ChannelValue {
	return State.SetValue(channelId, value)
}

func GetValue(channelId string) ChannelValue {
	return State.GetValue(channelId)
}

func RequestValue(channelId string) ChannelValue {
	return State.RequestValue(channelId)
}

func AllChannelStates() []ChannelState {
	states := make([]ChannelState, 0)
	for _, device := range AppConfig.Devices {
		if device.Enabled {
			for _, channel := range device.Channels {
				if channel.Enabled {
					states = append(states, ChannelState{channel, GetValue(channel.Name)})
				}
			}
		}
	}
	return states
}
