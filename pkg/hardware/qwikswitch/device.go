package qwikswitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	se "github.com/zdebeer99/zhome/pkg/stateengine"
)

type QwikSwitchDevice struct {
	id         string
	Connection string
	handler    se.EventHandler
	status     *se.Status
	running    bool
	channels   map[string]string
}

func New(id string, connection string) *QwikSwitchDevice {
	return &QwikSwitchDevice{
		id,
		connection,
		nil,
		se.NewStatus(),
		false,
		make(map[string]string),
	}
}

func (this *QwikSwitchDevice) ID() string {
	return this.id
}

func (this *QwikSwitchDevice) RegisterChannel(id string, address string, chType string) {
	this.channels[address] = id
}

func (this *QwikSwitchDevice) RegisterEventHandler(handler se.EventHandler) {
	this.handler = handler
}

func (this *QwikSwitchDevice) Status() *se.Status {
	return this.status
}

func (this *QwikSwitchDevice) Start() {
	this.running = true
	go this.worker()
}

func (this *QwikSwitchDevice) Stop() {
	this.running = false
}

//TODO: parse qs results and send the correct rsponse.
func (this *QwikSwitchDevice) SetValue(address string, value se.ValueMap) error {
	res, err := this.send(fmt.Sprintf("%s=%s", address, encodeValue(value.Value())))
	log.Println("SetValue ", res, err)
	return err
}

//TODO: parse qs results.
func (this *QwikSwitchDevice) GetValue(address string) (se.ValueMap, error) {
	res, err := this.send(fmt.Sprintf("%s?", address))
	log.Println("GetValue ", res, err)
	return res, err
}

func (this *QwikSwitchDevice) getConnection() string {
	return this.Connection
}

func (this *QwikSwitchDevice) buildUrl(cmd string) string {
	return this.getConnection() + "/" + cmd
}

func (this *QwikSwitchDevice) send(cmd string) (se.ValueMap, error) {
	resp, err := http.Get(this.buildUrl(cmd))
	if err != nil {
		return nil, fmt.Errorf("send() http.get error. %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("send() Read error. %s", err)
	}
	var data se.ValueMap
	json.Unmarshal(body, &data)
	log.Println("Body ", string(body))
	return data, nil
}

func (this *QwikSwitchDevice) worker() {
	log.Println("Qwikswitch worker started.")
	var safetyCount int
	for this.running {
		start := time.Now()
		this.listenOnce()
		//handle run away go thread.
		if (time.Since(start)) < (100 * time.Millisecond) {
			safetyCount = safetyCount + 1
			if safetyCount > 5 {
				log.Printf("qwikswitch.listen() Long pole returned in %v milliseconds to qwick! number of times %v", time.Since(start)*time.Millisecond, safetyCount)
				time.Sleep(10 * time.Second)
			}
			continue
		}
		safetyCount = 0
	}
	log.Println("Qwikswitch Listen Thread Ended.")
}

func (this *QwikSwitchDevice) listenOnce() {
	resp, err := this.send("&listen")
	if err != nil {
		log.Printf("listenOnce() http.get error. %s", err)
		time.Sleep(10 * time.Second)
		return
	}
	log.Println("Recieved Trigger ", resp)
	this.notify("recieved", resp["id"], resp)
	return
}

func (this *QwikSwitchDevice) notify(event string, chid string, value se.ValueMap) {
	if this.handler != nil {
		tevent := "switch" + chid
		this.handler(tevent, this.channels[chid], value)
	}
}

func encodeValue(value string) string {
	switch value {
	case "on":
		return "100"
	case "off":
		return "0"
	default:
		return value
	}
}
