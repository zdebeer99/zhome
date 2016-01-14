package stateengine

/*
This is a draft quick hack to log tempreture every 10 min. Must be updated to schedule jobs.
*/

import (
	"container/list"
	"time"
	//"github.com/zdebeer99/zhome/pkg/data"
)

func (this *StateEngine) startTaskWorker() {
	this.tickerTasks = time.NewTicker(300 * time.Millisecond)
	this.taskQueue = list.New()
	go func() {
		var element *list.Element
		for range this.tickerTasks.C {
			element = this.taskQueue.Front()
			for element != nil {
				e1 := element
				element = element.Next()
				t1 := e1.Value.(task)
				if time.Since(t1.targetTime) > (0 * time.Second) {
					this.taskQueue.Remove(e1)
					this.parseCommands(t1.commands)
				}
			}
		}
	}()
}

func (this *StateEngine) scheduleTask(target time.Time, commands []string) {
	for e := this.taskQueue.Front(); e != nil; e = e.Next() {
		t1 := e.Value.(task)
		if target.Before(t1.targetTime) {
			this.taskQueue.InsertBefore(task{time.Now(), target, commands}, e)
			return
		}
	}
	this.taskQueue.PushBack(task{time.Now(), target, commands})
}

// /*
// this is a temp function to record the temp every 10 min
// */
// func (this *StateEngine) startScheduler() {
// 	this.tickerScheduler = time.NewTicker(1 * time.Second)
// 	start := time.Now()
// 	for range this.tickerScheduler.C {
// 		if time.Now().Sub(start) > (10 * time.Minute) {
// 			start = time.Now()
// 			this.logTempreture()
// 		}
// 	}
// }
//
// type Measurement struct {
// 	SampleDate  time.Time
// 	ChannelType string
// 	ChannelID   string
// 	Value       ChannelValue
// }
//
// func (this *StateEngine) logTempreture() {
// 	db := data.DB("")
// 	defer db.Close()
// 	for id, ch := range this.channels {
// 		if ch.ChannelType == "dht22" {
// 			val1 := this.RequestValue(id)
// 			d1 := Measurement{
// 				SampleDate:  time.Now(),
// 				ChannelID:   id,
// 				ChannelType: ch.ChannelType,
// 				Value:       val1,
// 			}
// 			db.C("log_measurement").Insert(d1)
// 		}
// 	}
// }
