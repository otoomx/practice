package main

import (
	"fmt"
	"github.com/otoomx/practice/eventing/broker"
	"github.com/otoomx/practice/eventing/config"
	"log"
)

func main() {

	worker()

}

func worker() {
	log.Printf("Starting worker ...")
	broker, err := broker.New(config.Get())

	if err != nil {
		log.Println(err)
	} else {
		for i := 0; i < 5; i++ {
			go func(worker int){
			broker.CreateConsumer(func(msg []byte) {
				fmt.Printf("Consumer %d received message %s\n", worker, msg)
			})}(i)
		}

	}

	forever := make(chan bool)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
