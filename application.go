package scalyr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var _config Config

type Application interface {
	Event(sev Severity, attributes interface{})
}

type app struct {
	config Config
}

//
// event should accept Severity ( 0 - 6 ??, 3 being info:  and a message :  and 1 or more additional attributes)
// The "sev" (severity) field should range from 0 to 6, and identifies the importance of this event, using the
//  classic scale "finest, finer, fine, info, warning, error, fatal". This field is optional (defaults to 3 / info).
//
func (app *app) Event(sev Severity, attributes interface{}) {
	//Print(message)
	//Print(string(sev))

	// TODO:  test for scalyr api key being present...

	si := &scalyrSessionInfo{
		ServerType: "server type...",
		ServerId:   _APPLICATION_HOSTNAME,
	}

	se := &scalyrEvent{
		TS:    strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:  0,
		Sev:   int(sev),
		Attrs: attributes,
	}

	events := make([]scalyrEvent, 1)
	events[0] = *se

	ses := &scalyrEvents{
		Token:       _config.ScalyrEventWriteKey,
		Session:     _APPLICATION_PROCESS_ID,
		SessionInfo: *si,
		Events:      events,
	}

	json, _ := json.Marshal(ses)
	fmt.Println(string(json))

	url := _SCALYR_ADDEVENTS_ENDPOINT
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//
// returns a new config struct configured with values...
//
func NewApplication(config Config) Application {

	_config = config

	a := &app{}

	start()

	return a
}

func start() {
	ticker := time.NewTicker(time.Millisecond * 20000)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			//
			//
			// DO SOME WORK HERE!!!  push logs. etc...
			//
			//
			//
		}
	}()
}

/*
func Event(sev Severity, attributes interface{}) {

} */
