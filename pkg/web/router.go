//code generated from srcgen
package web

import (
	"github.com/gorilla/mux"
)

func RegisterApi(r *mux.Router) {
	// handlersHardware.go hardware comms
	r.HandleFunc("/setValue/{id}/{value}", handlerSetValue).Methods("GET")
	r.HandleFunc("/getValue/{id}", handlerGetValue).Methods("GET")
	r.HandleFunc("/channelStates", handlerChannelStates).Methods("GET")

	// handlersHardware.go Measurements
	//r.HandleFunc("/measurements/{id}", handlerMeasurements).Methods("GET")

}
