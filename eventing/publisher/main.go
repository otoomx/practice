package main

import (
	"github.com/otoomx/practice/eventing/broker"
	"github.com/otoomx/practice/eventing/config"
	"fmt"
	"log"
)

func main() {

	broker, _ := broker.New(config.Get())

	p, _ := broker.CreateProducer()

	for i:=0;i<300000;i++{
		p.Publish([]byte(fmt.Sprintf("Hello from %d", i)))
	}

	forever := make(chan bool)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
