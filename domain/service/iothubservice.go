package service

import (
	"caidc_auto_devicetwins/domain/utils"
	"crypto/tls"
	"fmt"
	"log"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConfig struct {
	client           mqtt.Client
	endpoint         string
	connectionString string
	uri              *url.URL
}

func (q *MQTTConfig) CreateDevice() string {
	return ""
}

func (q *MQTTConfig) PublishMessage(message string, deviceID string) bool {
	connected := q.connect(deviceID)
	if connected {
		if token := q.client.Publish(fmt.Sprintf("devices/%s/messages/events/", deviceID)+"$.ct=application%2Fjson&$.ce=utf-8",
			0, false, message); token.Wait() && token.Error() != nil {
			utils.FailOnError(token.Error(), "Error publishing message ")
		}
		return true
	}
	return false
}

func (q *MQTTConfig) connect(deviceID string) bool {
	opts := createClientOptions(deviceID, q.endpoint, q.connectionString)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		utils.FailOnError(token.Error(), "Error connecting to iothub")
	}
	log.Println("[MQTT] Connected")

	q.client = client
	return true
}

func createClientOptions(clientID string, endpoint string, sasToken string) *mqtt.ClientOptions {

	tlsConfig := tls.Config{InsecureSkipVerify: true}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcps://%s:8883", endpoint))
	opts.SetTLSConfig(&tlsConfig)
	opts.SetUsername(fmt.Sprintf("%s/%s/?api-version=2018-06-30", endpoint, clientID))
	opts.SetCleanSession(true)
	opts.SetPassword(sasToken)
	opts.SetClientID(clientID)
	return opts
}

func (q *MQTTConfig) Listen(topic string) {
	q.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
