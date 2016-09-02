package models

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/ryanuber/go-glob"
	"time"
)

type ResponseDelay struct {
	Path        *string             `json:"path"`
	Method      *string             `json:"method"`
	Destination *string             `json:"destination"`
	Scheme      *string             `json:"scheme"`
	Query       *string             `json:"query"`
	Body        *string             `json:"body"`
	Headers     map[string][]string `json:"headers"`
	Delay       int                 `json:"delay"`
}

type ResponseDelayPayload struct {
	Data *ResponseDelayList `json:"data"`
}

type ResponseDelayList []ResponseDelay

type ResponseDelays interface {
	Json() []byte
	GetDelay(request RequestDetails, webserver bool) *ResponseDelay
	Len() int
}

func (this *ResponseDelay) Execute() {
	// apply the delay - must be called from goroutine handling the request
	log.Info("Pausing before sending the response to simulate delays")
	time.Sleep(time.Duration(this.Delay) * time.Millisecond)
	log.Info("Response delay completed")
}

func (this *ResponseDelayList) GetDelay(request RequestDetails, webserver bool) *ResponseDelay {
	// iterate through the request templates, looking for template to match request
	for _, entry := range *this {
		if entry.Body != nil && !glob.Glob(*entry.Body, request.Body) {
			continue
		}

		if !webserver {
			if entry.Destination != nil && !glob.Glob(*entry.Destination, request.Destination) {
				continue
			}
		}
		if entry.Path != nil && !glob.Glob(*entry.Path, request.Path) {
			continue
		}
		if entry.Query != nil && !glob.Glob(*entry.Query, request.Query) {
			continue
		}
		if entry.Method != nil && !glob.Glob(*entry.Method, request.Method) {
			continue
		}
		if entry.Scheme != nil && !glob.Glob(*entry.Scheme, request.Scheme) {
			continue
		}

		// return the first template to match
		return &entry
	}
	return nil
}

func (this *ResponseDelayList) Json() []byte {
	resp := ResponseDelayPayload{
		Data: this,
	}
	b, _ := json.Marshal(resp)
	return b
}

func (this *ResponseDelayList) Len() int {
	list := []ResponseDelay{}
	if this != nil {
		list = append(list, *this...)
	}
	return len(list)
}
