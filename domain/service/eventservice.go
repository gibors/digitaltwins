package service

import (
	"caidc_auto_devicetwins/domain/factory"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type QueueConfig struct {
	Key        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	URL        string
	Name       string
	Token      string
}

func (q *QueueConfig) CreateQueConnection(queName string) {

	connString := fmt.Sprintf("amqp://:%s@%s/", q.Token, q.URL)
	log.Println(connString)
	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")
	q.Connection = conn

	ch, err := q.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	q.Channel = ch
	q.Name = queName
}

func (q *QueueConfig) CloseConnection() {
	if q.Channel != nil {
		defer q.Channel.Close()
	}
	if q.Connection != nil {
		defer q.Connection.Close()
	}
}

func (q *QueueConfig) PublishEventToRabbit(eventType string, message string) {
	q.CreateQueConnection(eventType)

	err := q.Channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
	q.CloseConnection()
}

func (q *QueueConfig) SubscribeToQueue(eventType string) {
	q.CreateQueConnection(eventType)
}

func (q *QueueConfig) TestConnection(queueName string) {
	q.CreateQueConnection(queueName)
	body := "Hello World!!!!!"
	err := q.Channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	q.CloseConnection()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateDeviceMessageEvent(dev device.Device, eventMessageType string) string {

	filePath := "./resources/" + eventMessageType + ".json"
	jsonFile, err := os.Open(filePath)

	failOnError(err, "Failed to open json file")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	newconnstr := string(byteValue)

	log.Println(newconnstr)

	return newconnstr
}

func GenerateNewConnectionData(d device.Device) string {
	connectionEvent := factory.CreateNewConnectionEvent(d)
	b, err := json.Marshal(connectionEvent)

	utils.FailOnError(err, "Failed to Marshall event object")

	return string(b)
}

func CreateTelemetryEvent(d device.Device) string {

	filePath := "./resources/telemetrymessage.json"
	jsonFile, err := os.Open(filePath)

	failOnError(err, "Failed to open json file")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	telemetryEvent := string(byteValue)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{CreatedTime}", utils.GenerateEventTimeStamp())
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SystemGUID}", d.SystemGUID)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SystemType}", d.SystemType)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SerialNumber}", d.SerialNumber)

	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.repcap.value}", string(generateRandomValue(5550, 5570)))  // reported capacity
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.repsoc.value}", string(generateRandomValue(10, 100)))     // reported state of charge
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.soh.value}", string(generateRandomValue(0, 100)))         // battery age value
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcap.value}", string(generateRandomValue(5520, 5590))) //mAh full capacity of battery at pressent
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.tte.value}", string(generateRandomValue(330000, 369090))) // tte
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.designcap.value}", string(generateRandomValue(5866, 5866)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxtemp.value}", string(generateRandomValue(31, 31)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.cyclecount.value}", string(generateRandomValue(30, 100)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.timerh.value}", string(generateRandomValue(19066400, 19086400)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.mintemp.value}", string(generateRandomValue(22, 22)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.minvolt.value}", string(generateRandomValue(3, 3)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxvolt.value}", string(generateRandomValue(4, 4)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxcurrent.value}", string(generateRandomValue(2, 4)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.mincurrent.value}", string(generateRandomValue(-1, 1)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.ttf.value}", string(generateRandomValue(5550, 5570)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcapnom.value}", string(generateRandomValue(366634, 369634)))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcaprep.value}", string(generateRandomValue(5548, 5588)))

	return telemetryEvent
}

func generateRandomValue(start int, end int) int {
	if end < start {
		log.Fatal("start should be less than end ")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start+1) + start
}