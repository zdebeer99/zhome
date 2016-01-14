package stateengine

import (
	"log"
)

func (this *StateEngine) eventHandler(eventName string, id string, arg interface{}) {
	log.Println("eventHandler Recieved Event.", eventName, len(this.triggers))
	for _, trigger := range this.triggers {
		if trigger.EventName == eventName {
			log.Println("eventHandler Found Trigger.", eventName)
			this.parseCommands(trigger.Command)
		}
	}
}
