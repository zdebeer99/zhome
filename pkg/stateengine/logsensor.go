/*
This code will need review, this section is a quick and nasty lsensor loggin solution.
*/
package stateengine

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

func (this *StateEngine) cmdLogSensor(addr string) error {
	chinfo := FindChannel(addr)
	if chinfo == nil {
		return fmt.Errorf("logSensor: Could not find channel '%s'.", addr)
	}
	value := this.RequestValue(addr)
	session, err := mgo.Dial(AppConfig.Database)
	if err != nil {
		return fmt.Errorf("Failed to connect to database. %s", err)
	}
	defer session.Close()

	var m = make(map[string]interface{})
	m["sampletime"] = time.Now()
	m["channel"] = addr
	m["info"] = chinfo
	m["statusCode"] = value.StatusCode
	if value.StatusCode != 0 {
		m["statusText"] = value.StatusText
	}
	var mdata = make(map[string]interface{})
	for k, v := range value.Data {
		v1, err := strconv.ParseFloat(v, 64)
		if err != nil {
			mdata[k] = v
		} else {
			mdata[k] = v1
		}
	}
	m["values"] = mdata

	db := session.DB("")
	err = db.C("log_sensor").Insert(m)
	if err != nil {
		return fmt.Errorf("Failed to Insert sensor data into database. %s", err)
	}
	return nil
}
