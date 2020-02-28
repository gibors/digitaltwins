package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MESSAGE = `{
	"CloudPlatformEvent": {
	  "CreatedTime": "2020-02-19T01:43:51-06:00",
	  "Id": "09eded4f-ccac-4b9e-a742-1dc5132541c6",
	  "CreatorId": "0f91a940-8567-4eb7-b8f5-a1f942b3fd99",
	  "CreatorType": "CloudPlatformSystem",
	  "GeneratorId": null,
	  "GeneratorType": "CloudPlatformTenant",
	  "TargetId": "0f91a940-8567-4eb7-b8f5-a1f942b3fd99",
	  "TargetType": "CloudPlatformTenant",
	  "TargetContext": null,
	  "Body": {
		"Type": "TextualBody",
		"Value": "{\"SystemGuid\":\"eeaa9280-849c-4c9f-b52c-fcb3952ce21b\",\"HistorySamples\":[{\"ItemName\":\"CT6003VR2FVES7.battery.repcap\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.315+00:00\",\"Value\":\"5575\"},{\"ItemName\":\"CT6003VR2FVES7.battery.repsoc\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.34+00:00\",\"Value\":\"43\"},{\"ItemName\":\"CT6003VR2FVES7.battery.soh\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.36+00:00\",\"Value\":\"96\"},{\"ItemName\":\"CT6003VR2FVES7.battery.fullcap\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.378+00:00\",\"Value\":\"5582\"},{\"ItemName\":\"CT6003VR2FVES7.battery.tte\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.395+00:00\",\"Value\":\"353564\"},{\"ItemName\":\"CT6003VR2FVES7.battery.designcap\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.411+00:00\",\"Value\":\"5866\"},{\"ItemName\":\"CT6003VR2FVES7.battery.maxtemp\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.429+00:00\",\"Value\":\"31\"},{\"ItemName\":\"CT6003VR2FVES7.battery.mintemp\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.447+00:00\",\"Value\":\"22\"},{\"ItemName\":\"CT6003VR2FVES7.battery.maxvolt\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.465+00:00\",\"Value\":\"4\"},{\"ItemName\":\"CT6003VR2FVES7.battery.minvolt\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.48+00:00\",\"Value\":\"3\"},{\"ItemName\":\"CT6003VR2FVES7.battery.maxcurrent\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.505+00:00\",\"Value\":\"3\"},{\"ItemName\":\"CT6003VR2FVES7.battery.mincurrent\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.532+00:00\",\"Value\":\"0\"},{\"ItemName\":\"CT6003VR2FVES7.battery.ttf\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.552+00:00\",\"Value\":\"5563\"},{\"ItemName\":\"CT6003VR2FVES7.battery.fullcapnom\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.57+00:00\",\"Value\":\"368950\"},{\"ItemName\":\"CT6003VR2FVES7.battery.fullcaprep\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.588+00:00\",\"Value\":\"5588\"},{\"ItemName\":\"CT6003VR2FVES7.battery.timerh\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.607+00:00\",\"Value\":\"19068696\"},{\"ItemName\":\"CT6003VR2FVES7.battery.cyclecount\",\"Quality\":\"Good\",\"Time\":\"2018-12-05T03:58:56.626+00:00\",\"Value\":\"92\"}]}",
		"Format": "application/json"
	  },
	  "BodyProperties": [
		{
		  "Key": "SystemType",
		  "Value": ""
		},
		{
		  "Key": "SystemGuid",
		  "Value": "eeaa9280-849c-4c9f-b52c-fcb3952ce21b"
		}
	  ],
	  "EventType": "DataChange.Update"
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

func TestPublishMessage(t *testing.T) {
	config := MQTTConfig{}
	config.endpoint = "caidc-dev-iothub.azure-devices.net"
	config.connectionString = "SharedAccessSignature sr=caidc-dev-iothub.azure-devices.net&sig=VU2367nY%2BFkespWYsiF6IWCW2wDnrAGGjTEKVTO4WDI%3D&se=1614125551&skn=iothubowner"
	deviceID := "eeaa9280-849c-4c9f-b52c-fcb3952ce21b"
	config.PublishMessage(MESSAGE, deviceID)
}
