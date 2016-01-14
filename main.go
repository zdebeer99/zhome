package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/zdebeer99/weblib/handler"
	"github.com/zdebeer99/zhome/pkg/config"
	"github.com/zdebeer99/zhome/pkg/controllers"
	"github.com/zdebeer99/zhome/pkg/hardware"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
	"log"
)

var appConfig = config.Load()

//***
func main() {
	log.Println("Connecting to ", appConfig.InfluxServer)

	load(appConfig)

	r := mux.NewRouter()
	r.PathPrefix("/html/").Handler(handler.FileServer("./static/"))
	r.PathPrefix("/css/").Handler(handler.FileServer("./static/"))
	r.PathPrefix("/js/").Handler(handler.FileServer("./static"))

	r.Handle("/charts", handler.File("./static/html/charts/index.html"))
	//r.Handle("/config", handler.File("./static/html/config/index.html"))
	r.Handle("/", handler.File("./static/html/home/index.html"))
	controllers.RegisterApi(r.PathPrefix("/api").Subrouter())

	n := negroni.Classic()
	n.UseHandler(r)
	log.Println("Run web service. on ", appConfig.BindAddress)
	n.Run(appConfig.BindAddress)
}

// load stateengine setup data.
func load(cfg *config.Config) {
	se.AppConfig = cfg
	state := se.State
	//Register devices
	for _, device := range cfg.Devices {
		state.RegisterDevice(device.Name, hardware.NewDeviceComm(device.DeviceType, device.Name, device.Connection))
		//Register Channels
		for _, ch := range device.Channels {
			state.RegisterChannel(ch.Name, device.Name, ch.Address, ch.ChannelType)
		}
	}

	//Register Events
	for _, trigger := range cfg.Triggers {
		state.AddTrigger(se.Trigger(trigger))
	}

	//Register Schedules
	for _, schedule := range cfg.Scheduler {
		state.AddSchedule(schedule.Name, schedule.When, schedule.Command)
	}
	log.Println("Config Loaded.")
	se.Start()
}
