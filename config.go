package scalyr

import "time"

const _SCALYR_BASE_ENDPOINT = "https://www.scalyr.com/"
const _SCALYR_ADDEVENTS_ENDPOINT = _SCALYR_BASE_ENDPOINT + "addEvents"

var _APPLICATION_PROCESS_ID = UUID()
var _APPLICATION_HOSTNAME = HostName()

// TODO rename this var name !!

var ticker = time.NewTicker(time.Millisecond * 5000) // every 5 seconds

type Config struct {

	//
	// should the agent write out local logs for debugging purposes.
	// useful for when you are using the agent for the first time or
	// for tracking down an issue....
	//
	LocalLog bool

	//
	// for now, just requiring the scalyr event write key...
	// in the future might add the config keys and read keys...
	//
	ScalyrEventWriteKey string

	//
	// Tags these logs as coming from central place
	// e.g. AppServer1, web-app-2, db-server-99...
	//
	LogTag string
}

//
// returns a new config struct configured with values...
//
func NewConfig(localLog bool, logTag string, scalyrEventWriteKey string) Config {

	c := Config{}

	c.LocalLog = localLog
	//
	// TODO:  test for scalyr api key being present...
	//
	c.ScalyrEventWriteKey = scalyrEventWriteKey
	c.LogTag = logTag

	return c
}
