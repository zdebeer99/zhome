package web

import (
	"github.com/gorilla/mux"
	"github.com/zdebeer99/weblib"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
	"net/http"
)

func handlerSetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := se.SetValue(vars["id"], vars["value"])
	weblib.WriteResponse(w, data, nil)
}

func handlerGetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := se.RequestValue(vars["id"])
	weblib.WriteResponse(w, data, nil)
}

// deviceList
// Function generated from srcgen
func handlerChannelStates(wr http.ResponseWriter, req *http.Request) {
	weblib.WriteResponse(wr, se.AllChannelStates(), nil)
}

// func handlerMeasurements(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	channelid := vars["id"]
// 	var samples []data.M
// 	err := db.C(data.C_MEASUREMENT).
// 		Find(data.M{"channelid": channelid}).
// 		Sort("-sampledate").
// 		Limit(1000).
// 		All(&samples)
// 	weblib.WriteResponse(w, samples, err)
// }
