package main

import (
	"github.com/otoomx/practice/eventing/config"
	"github.com/streadway/amqp"
	"log"
	"time"

)

//internal amqp connection wrapper
//composes connection and channel
type amqpConn struct {
	*amqp.Connection
	*amqp.Channel
	config           *config.Config
	rabbitCloseError chan *amqp.Error
}

func reconnect(conn *amqpConn) *amqpConn {
	for {
		var (
			c  *amqp.Connection
			ch *amqp.Channel
		)
		c, err := amqp.Dial(conn.config.Broker)
		if err == nil {
			ch, err = c.Channel()
		}
		log.Printf("Trying to reconnect to RabbitMQ at %s\n", conn.config.Broker)
		if err == nil {
			conn.Connection = c
			conn.Channel = ch
			return conn
		}
		log.Println(err)
		time.Sleep(500 * time.Millisecond)
	}
}

func connect(conn *amqpConn) {
	var rabbitErr *amqp.Error
	for {
		rabbitErr = <-conn.rabbitCloseError
		if rabbitErr != nil {
			log.Println("Rabbit Close error")
			reconnect(conn)
			conn.rabbitCloseError = make(chan *amqp.Error)
			conn.Connection.NotifyClose(conn.rabbitCloseError)
		}

		q, err := conn.QueueDeclare(
			"test", //queue name
			true,  //message should be persistant
			false, //auto delete no
			false, // exclusive is false, set to true will remove persistance
			false, //no wait is false because we want to know if rabbit cant create queue
			nil)

		err = conn.QueueBind(q.Name, conn.config.BindingKey, conn.config.Exchange, false, nil)
		log.Println(err)

		replies, err := conn.Consume(
			q.Name,             // queue
			q.Name+time.Now().String(), // consumer
			false,              // auto-ack
			false,               // exclusive, there shouldnt be any other consumer
			false,              // no-local
			false,              // no-wait
			nil,                // args
		)

		log.Printf("consume")

		go func() {
			for d := range replies {
				log.Println(d.Body)
				d.Ack(false)
				//message <- &Msg{Body: d.Body, Transport: d}
			}
		}()
	}
}

func New(config *config.Config) *amqpConn {

	//allocate a new session
	session := new(amqpConn)
	session.config = config
	// create closeErroe channel
	session.rabbitCloseError = make(chan *amqp.Error)

	go connect(session)
	//send close error to channel
	session.rabbitCloseError <- amqp.ErrClosed

	return session
}

func main() {

	New(config.Get())
	forever := make(chan bool)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
