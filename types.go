package scalyr

type Severity int

type scalyrEvent struct {
	TS    string      `json:"ts"`
	Type  int         `json:"type"`
	Sev   int         `json:"sev"`
	Attrs interface{} `json:"attrs"`
}

type scalyrSessionInfo struct {
	ServerType string `json:"serverType"`
	ServerId   string `json:"serverId"`
}

type scalyrEventsWrapper struct {
	Token       string            `json:"token"`
	Session     string            `json:"session"`
	SessionInfo scalyrSessionInfo `json:"sessionInfo"`
	Events      []scalyrEvent     `json:"events"`
}

const (
	Debug Severity = 2 + iota
	Info
	Warning
	Error
	Fatal
)

var severityLevels = [...]string{
	"Debug",
	"Info",
	"Warning",
	"Error",
	"Fatal",
}

// String returns the English name of the month ("January", "February", ...).
func (sev Severity) String() string { return severityLevels[sev-2] }

func TestSeverityLevel(sev Severity) {

	if sev < 2 || sev > 6 {
		panic("Severity value was out of range!")
	}

}
