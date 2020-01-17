package utils

import (
	"log"
	"time"
)

func GenerateEventTimeStamp() string {
	tm := time.Now()
	log.Printf("Current time: %s", tm.String())
	return tm.Format(time.RFC3339)
}
