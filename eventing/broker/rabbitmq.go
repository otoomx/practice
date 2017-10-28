package broker

import (
	"github.com/otoomx/practice/eventing/config"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// - Broker Implementation
type amqpBroker struct {
	amqp.Connection
	amqp.Channel
	closeError chan *amqp.Error
	config *config.Config
}

func (broker *amqpBroker) CreateConsumer(fn CallBack) (MsgConsumer, error) {
	c := broker.config

	//ch.Qos(1,1,true)
	log.Printf("Creating new consumer for exchange: %s routing key: %s\n", c.Exchange, c.BindingKey)
	msgQueue, err := broker.QueueDeclare(
		"test", //queue name
		true,  //message should be persistant
		false, //auto delete no
		false, // exclusive is false, set to true will remove persistance
		false, //no wait is false because we want to know if rabbit cant create queue
		nil)
	log.Printf("Binding to queue for : %s routing key: %s\n", c.Exchange, c.BindingKey)

	err = broker.QueueBind(msgQueue.Name, c.BindingKey, c.Exchange, false, nil)
	log.Println(err)




	log.Printf("Starting comsumer : %s routing key: %s\n", c.Exchange, c.BindingKey)
	replies, err := broker.Consume(
		msgQueue.Name,             // queue
		msgQueue.Name+time.Now().String(), // consumer
		false,              // auto-ack
		false,               // exclusive, there shouldnt be any other consumer
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)

	log.Printf("consume")

	go func() {
		for d := range replies {
			fn(d.Body)
			d.Ack(false)
			//message <- &Msg{Body: d.Body, Transport: d}
		}
	}()

	return &amqpConsumer{broker, nil}, err
}

func (broker *amqpBroker) CreateProducer() (MsgProducer, error) {
	c := broker.config

	log.Printf("Creating new producer for exchange: %s routing key: %s\n", c.Exchange, c.BindingKey)

	err := broker.ExchangeDeclare(
		c.Exchange,     // name
		"direct", // type
		true,     // durable
		false,          // auto-deleted
		false,          // internal
		false,          // noWait
		nil,            // arguments

	)
	if err != nil {
		return nil, err
	}
	return &amqpProducer{broker}, nil
}

func (broker *amqpBroker) Close() {

}

// - Producer Implementation

type amqpProducer struct {
	*amqpBroker
}

func (p amqpProducer) Publish(msg []byte) error {

	log.Println("publisjing message")
	err := p.Channel.Publish(
		p.config.Exchange,   // exchange
		p.config.BindingKey, // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	return err
}

// - Consumer Implementation
type amqpConsumer struct {
	*amqpBroker
	msgCh chan *Msg
}

func (c amqpConsumer) OnMessage(msg []byte) {

}

//New returns new Msg broker based on the configuration settings
func New(config *config.Config) (MsgBroker, error) {

	//create connection
	log.Printf("Connecting to broker %s", config.Broker)
	conn, err := amqp.Dial(config.Broker)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	//create channel


	if err != nil {
		return nil, err
	}

	b := &amqpBroker{*conn, *ch, make(chan *amqp.Error), config}
	log.Println("Returning broker")
	return b, nil
}
