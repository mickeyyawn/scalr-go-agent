package scalyr

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var _config Config
var _scalyrEventsWrapper scalyrEventsWrapper
var _scalyrEvents []scalyrEvent

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

	se := &scalyrEvent{
		TS:    strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:  0,
		Sev:   int(sev),
		Attrs: attributes,
	}

	if _scalyrEvents == nil {

		_scalyrEvents = make([]scalyrEvent, 1)
		_scalyrEvents[0] = *se

	} else {
		_scalyrEvents = append(_scalyrEvents, *se)
	}

}

//
// returns a new config struct configured with values...
//
func NewApplication(config Config) Application {

	_config = config

	a := &app{}

	initialize()

	start()

	return a
}

func initialize() {
	si := &scalyrSessionInfo{
		ServerType: "server type here TODO replace this...",
		ServerId:   _APPLICATION_HOSTNAME,
	}

	_scalyrEventsWrapper = scalyrEventsWrapper{
		Token:       _config.ScalyrEventWriteKey,
		Session:     _APPLICATION_PROCESS_ID,
		SessionInfo: *si,
		Events:      nil,
	}

	Print("Events wrapper created: ", _scalyrEventsWrapper)
}

func start() {

	go func() {
		for range ticker.C {
			flush()
		}
	}()

}

func flush() {
	Print("Flushing events NOW. ", nil)

	_scalyrEventsWrapper.Events = _scalyrEvents

	json, _ := json.Marshal(_scalyrEventsWrapper)

	Print("Events being sent to Scalyr: ", string(json))
	//fmt.Println(string(json))

	url := _SCALYR_ADDEVENTS_ENDPOINT

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	Print("Status that came back from Scalyr: ", resp.Status)

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
