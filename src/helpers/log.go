package helpers

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// ELK is a
type ELK struct {
	URILog      string
	RequestLog  interface{}
	ResponseLog interface{}
	StatusLog   string
}

// ReqRes is log generator for server to elastic and kibana
func (o ELK) ReqRes() *logrus.Entry {
	var response interface{}
	o.StatusLog = "Success"

	if o.ResponseLog == response {
		o.StatusLog = "Failed"
	}

	return log.WithFields(logrus.Fields{
		"uri":      o.URILog,
		"request":  o.RequestLog,
		"response": o.ResponseLog,
		"status":   o.StatusLog,
	})
}

// LogInfo is a
func (o ELK) LogInfo(msg string) {
	o.ReqRes().Info(msg)
}

// LogWarn is a
func (o ELK) LogWarn(msg string) {
	o.ReqRes().Warn(msg)
}

// LogDebug merupakan fungsi untuk logging format text INFO[2021-03-11 09:08:11] [202103190001] - msg
// parameter refnum optional
func (o ELK) LogDebug(msg string, refnum ...string) {
	setText()
	log.Infof("%s - %v", refnum, msg)
	setJSON()
}

func setText() {
	formatter := new(log.TextFormatter)
	formatter.DisableQuote = true
	formatter.DisableSorting = true
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.ForceColors = true
	formatter.DisableLevelTruncation = true
	log.SetFormatter(formatter)
}

func setJSON() {
	formatter := new(log.JSONFormatter)
	// formatter.PrettyPrint = true
	log.SetFormatter(formatter)
}
