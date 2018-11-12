package rbot

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func CreateQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func CreateConsumeChannel(ch *amqp.Channel, name string) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func CreateRpcBase(connection *amqp.Connection) (*amqp.Channel, amqp.Queue, <-chan amqp.Delivery, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, amqp.Queue{}, make(chan amqp.Delivery), NewErrorRemoteBot(FailedOpenChannel, err)
	}

	q, err := CreateQueue(ch)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, make(chan amqp.Delivery), NewErrorRemoteBot(FailedDeclareQueue, err)
	}

	msgs, err := CreateConsumeChannel(ch, q.Name)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, make(chan amqp.Delivery), NewErrorRemoteBot(FailedMessageConsume, err)
	}

	return ch, q, msgs, nil
}

func Publish(ch *amqp.Channel, requestMessage *RequestMessage, name string, request []byte) error {
	return ch.Publish(
		"",         // exchange
		RoutingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       name,
			Body:          request,
		})
}

func PublishWithTimeout(ch *amqp.Channel, requestMessage *RequestMessage, name string, request []byte, ticker *time.Ticker) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- Publish(ch, requestMessage, name, request)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return NewErrorRemoteBot(FailedMessagePublish, err)
		}
	case <-ticker.C:
		return NewErrorRemoteBot(FailedMessagePublish, fmt.Errorf("timeout"))
	}

	return nil
}

func ConsumeWithTimeout(correlationId string, msgs <-chan amqp.Delivery, ticker *time.Ticker) (*ResponseMessage, error) {
	var response ResponseMessage
	errChan := make(chan error, 1)

	go func() {
		for d := range msgs {
			if correlationId == d.CorrelationId {
				err := json.Unmarshal(d.Body, &response)
				errChan <- err
				break
			}
		}
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, NewErrorRemoteBot(FailedConvertBodyResponse, err)
		}
	case <-ticker.C:
		return nil, NewErrorRemoteBot(FailedMessageConsume, fmt.Errorf("timeout"))
	}

	return &response, nil
}

func RpcWithTimeout(ch *amqp.Channel, q amqp.Queue, msgs <-chan amqp.Delivery, requestMessage *RequestMessage, ticker *time.Ticker) (*ResponseMessage, error) {
	request, err := json.Marshal(*requestMessage)
	if err != nil {
		panic(err)
	}

	remoteBotErr := PublishWithTimeout(ch, requestMessage, q.Name, request, ticker)
	if remoteBotErr != nil {
		return nil, remoteBotErr
	}

	response, remoteBotErr := ConsumeWithTimeout(requestMessage.CorrelationId, msgs, ticker)
	if remoteBotErr != nil {
		return nil, remoteBotErr
	}

	return response, nil
}
