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
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type Broker interface {
	Connect() bool
	Publish() bool
}
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
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	q.Connection = conn

	ch, err := q.Connection.Channel()
	utils.FailOnError(err, "Failed to open a channel")
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
	utils.FailOnError(err, "Failed to publish a message")
	q.CloseConnection()
}

func (q *QueueConfig) SubscribeToQueue(eventType string) {
	q.CreateQueConnection(eventType)
}

func CreateDeviceMessageEvent(dev device.Device, eventMessageType string) string {

	filePath := "./resources/" + eventMessageType + ".json"
	jsonFile, err := os.Open(filePath)

	utils.FailOnError(err, "Failed to open json file")

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

	utils.FailOnError(err, "Failed to open json file")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	telemetryEvent := string(byteValue)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{CreatedTime}", utils.GenerateEventTimeStamp())
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SystemGUID}", d.SystemGUID)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SystemType}", d.SystemType)
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{SerialNumber}", d.SerialNumber)

	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.repcap.value}", strconv.FormatInt(generateRandomValue(5550, 5570), 10))  // reported capacity
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.repsoc.value}", strconv.FormatInt(generateRandomValue(10, 100), 10))     // reported state of charge
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.soh.value}", strconv.FormatInt(generateRandomValue(0, 100), 10))         // battery age value
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcap.value}", strconv.FormatInt(generateRandomValue(5520, 5590), 10)) //mAh full capacity of battery at pressent
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.tte.value}", strconv.FormatInt(generateRandomValue(330000, 369090), 10)) // tte
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.designcap.value}", strconv.FormatInt(generateRandomValue(5866, 5866), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxtemp.value}", strconv.FormatInt(generateRandomValue(31, 31), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.cyclecount.value}", strconv.FormatInt(generateRandomValue(30, 100), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.timerh.value}", strconv.FormatInt(generateRandomValue(19066400, 19086400), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.mintemp.value}", strconv.FormatInt(generateRandomValue(22, 22), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.minvolt.value}", strconv.FormatInt(generateRandomValue(3, 3), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxvolt.value}", strconv.FormatInt(generateRandomValue(4, 4), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.maxcurrent.value}", strconv.FormatInt(generateRandomValue(2, 4), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.mincurrent.value}", strconv.FormatInt(generateRandomValue(-1, 1), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.ttf.value}", strconv.FormatInt(generateRandomValue(5550, 5570), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcapnom.value}", strconv.FormatInt(generateRandomValue(366634, 369634), 10))
	telemetryEvent = strings.ReplaceAll(telemetryEvent, "{battery.fullcaprep.value}", strconv.FormatInt(generateRandomValue(5548, 5588), 10))

	return telemetryEvent
}

func generateRandomValue(start int, end int) int64 {
	if end < start {
		log.Fatal("start should be less than end ")
	}
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(end-start+1) + start
	return int64(number)
}
