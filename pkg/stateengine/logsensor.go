/*
This code will need review, this section is a quick and nasty lsensor loggin solution.
*/
package stateengine

import (
	"github.com/influxdb/influxdb/client/v2"
	"log"
	"time"
)

func logSensor(addr string, value ValueMap) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: AppConfig.InfluxServer})
	if err != nil {
		log.Println("Failed to write sensor data. ", err)
		return
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "home",
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{
		"location":   "living room",
		"sensor":     "temparature",
		"sensorType": "dht22",
		"sensorId":   "ks01",
	}
	fields := map[string]interface{}{
		"temperature": value["temp"],
		"humidity":    value["hum"],
	}
	pt, err := client.NewPoint(addr, tags, fields, time.Now())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}
