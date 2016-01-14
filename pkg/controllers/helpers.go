package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func getId(req *http.Request) string {
	vars := mux.Vars(req)
	return vars["id"]
}
