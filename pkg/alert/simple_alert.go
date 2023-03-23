package alert

import (
	"fmt"
)

const (
	criticalLevel = "CRITICAL"
	highLevel     = "HIGH"
	mediumLevel   = "MEDIUM"
	lowLevel      = "LOW"
)

var makeSimpleCriticalAlert = newSimplePayload(criticalLevel)
var makeSimpleHighAlert = newSimplePayload(highLevel)
var makeSimpleMediumAlert = newSimplePayload(mediumLevel)
var makeSimpleLowAlert = newSimplePayload(lowLevel)

func newSimplePayload(level string) func(Data) SimplePayload {
	return func(ad Data) SimplePayload {
		return SimplePayload{
			Title:   ad.Title,
			Message: ad.Message,
			Level:   level,
		}
	}
}

type SimplePayload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

func (h SimplePayload) String() string {
	return fmt.Sprintf("%s - %s", h.Title, h.Message)
}
