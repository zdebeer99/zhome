package stateengine

import (
	"fmt"
	"log"
	"strings"
)

func (this *StateEngine) parseCommands(commands []string) {
	for _, cmd := range commands {
		this.parseCommand(cmd)
	}
}

func (this *StateEngine) parseCommand(command string) {
	steps := strings.Split(command, "|")
	for _, cmdstr := range steps {
		cmdarg := strings.Split(cmdstr, " ")
		if len(cmdarg) > 0 {
			tcmd := strings.ToLower(cmdarg[0])
			var targ []string
			if len(cmdarg) > 1 {
				targ = cmdarg[1:]
			}
			err := this.execCommand(tcmd, targ)
			if err != nil {
				log.Printf("Error executing command '%s' %s", command, err)
			}
		} else {
			//TODO: Remove this log.
			log.Println("parseCommand EmptyCommand", command)
		}
	}
}

func (this *StateEngine) execCommand(cmd string, args []string) error {
	switch cmd {
	case "toggle":
		if len(args) < 1 {
			return fmt.Errorf("'toggle' command expects 1 argument containning a channel id")
		}
		return this.cmdToggel(args[0])
	case "on":
		if len(args) < 1 {
			return fmt.Errorf("'on' command expects 1 argument containning a channel id")
		}
		return this.cmdSet(args[0], "on")
	case "off":
		if len(args) < 1 {
			return fmt.Errorf("'off' command expects 1 argument containning a channel id")
		}
		return this.cmdSet(args[0], "off")
	case "set":
		if len(args) < 2 {
			return fmt.Errorf("'set' command expects 2 argument containning a channel id and a value")
		}
		return this.cmdSet(args[0], args[1])
	case "read":
		return fmt.Errorf("'read' not implemented yet")
	case "logsensor":
		if len(args) < 1 {
			return fmt.Errorf("'logSensor' command expects 1 argument containning a channel id.")
		}
		return this.cmdLogSensor(args[0])
	}
	return nil
}

func (this *StateEngine) cmdToggel(channelId string) error {
	val := this.GetValue(channelId)
	var r1 ChannelValue
	if val.StatusCode == 0 {
		if val.Data.Value() == "off" {
			r1 = this.SetValue(channelId, "on")
		} else {
			r1 = this.SetValue(channelId, "off")
		}
	} else {
		return fmt.Errorf("toggle command recieved error from channel '%s'. %s", channelId, val.StatusText)
	}
	if r1.StatusCode > 0 {
		return fmt.Errorf("toggle command could not set value on channel '%s'. %s", channelId, val.StatusText)
	}
	return nil
}

func (this *StateEngine) cmdSet(channelId, value string) error {
	val := this.SetValue(channelId, value)
	if val.StatusCode != 0 {
		return fmt.Errorf("Set command recieved error from channel '%s'. %s", channelId, val.StatusText)
	}
	return nil
}
