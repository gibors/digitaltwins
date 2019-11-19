package service

import (
	device "caidc_auto_devicetwins/domain/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type QueueConfig struct {
	Key        *string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	URL        *string
	Name       string
	Token      string
}

func (q *QueueConfig) CreateQueConnection(queName string) {

	connString := fmt.Sprintf("amqp://:%s@%s/", q.Token, *q.URL)
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

func (q *QueueConfig) PublishNewConnectionEvent(eventType string, message string) {
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

func GetNewConnectionEvent(dev device.Device) string {
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/resources/newconnectionmessage.json")

	failOnError(err, "Failed to open json file")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	newconnstr := string(byteValue)

	log.Println(newconnstr)

	return newconnstr
}
