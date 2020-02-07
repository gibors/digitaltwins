package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MESSAGE = `{
	"CloudPlatformEvent": {
	  "CreatedTime": "2018-11-22T09:17:36.803+05:30",
	  "Id": "ea4c5f57-962f-4336-a3a9-6e312bc1a48c",
	  "CreatorId": null,
	  "CreatorType": "EBICC R1",
	  "GeneratorId": null,
	  "GeneratorType": "CN8017340D8241",
	  "TargetId": "16:79:50:B2:6F:86",
	  "TargetType": null,
	  "TargetContext": null,
	  "Body": null,
	  "BodyProperties": [
		{
		  "Key": "OP_DESC",
		  "Value": ""
		},
		{
		  "Key": "DETAILS",
		  "Value": "{\"DeviceId\":\"CN8017340D8241\",\"RequestId\":\"62c01a61-b1ae-4e8a-9e6e-140e77d7bdc2\",\"ScheduledVersion\":\"5.3.0.41\"}"
		},
		{
		  "Key": "EventGeneratedTime",
		  "Value": "2018-11-22T03:47:36.802+00:00"
		},
		{
		  "Key": "OP_STATUS",
		  "Value": "FAILED"
		},
		{
		  "Key": "SystemType",
		  "Value": "opintel"
		},
		{
		  "Key": "SystemGuid",
		  "Value": "81dfa3c6-03ea-4ae2-8c4d-4dc25bee37a5"
		}
	  ],
	  "EventType": "NOTIFICATION"
	},
	"AnnotationStreamIds": ","
  }`

func TestTestConnection(t *testing.T) {
	queueConfig := QueueConfig{}
	queueConfig.URL = "40.77.30.88:5672"
	queueConfig.PublishEventToRabbit("notifications", MESSAGE)
	log.Println(queueConfig)
	assert.NotNil(t, queueConfig.Connection)
}

func TestSuscribeEvent(t *testing.T) {
	// queueConfig := QueueConfig{}
}
