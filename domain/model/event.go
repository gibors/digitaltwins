package device

import "time"

type CloudPlatformEvent struct {
	CreatedTime    time.Time      `json:"createdTime"`
	ID             *string        `json:"Id"`
	CreatorID      *string        `json:"CreatorId"`
	CreatorType    *string        `json:"CreatorType"`
	GeneratorID    *string        `json:"GeneratorId"`
	GeneratorType  *string        `json:"GeneratorType"`
	TargetID       *string        `json:"TargetId"`
	TargetType     *string        `json:"TargetType"`
	TargetContext  *string        `json:"TargetContext"`
	BodyMessage    *Body          `json:"Body"`
	BodyProperties []BodyProperty `json:"BodyProperties"`
	EventType      *string        `json:"EventType"`
}

type Body struct {
	Value  string
	Type   string
	Format string
}

type BodyProperty struct {
	Key   string
	Value string
}
