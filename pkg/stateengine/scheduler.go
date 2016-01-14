package stateengine

/*
the scheduler keep track of scheduled commands.
*/

/*
Look at scheduler libraries
* https://github.com/gorhill/cronexpr
  Current choice.

* https://godoc.org/github.com/robfig/cron
  Looks Promising
* https://github.com/jasonlvhit/gocron/blob/master/gocron.go
  No String parsing
*/

import (
	"github.com/robfig/cron"
)

type cronJob struct {
	se       *StateEngine
	schedule Schedule
}

func (this *cronJob) Run() {
	this.se.parseCommands(this.schedule.Commands)
}

func (this *StateEngine) startScheduler() {
	this.Scheduler = cron.New()
	for _, sched := range this.schedules {
		this.Scheduler.AddJob(sched.When, &cronJob{this, sched})
	}
	this.Scheduler.Start()
}
