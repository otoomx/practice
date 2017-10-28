// broker package is an abstraction of a message broker.  It tries to hide the implementation details of the actual
// vendor specific msg broker details
package broker


type CallBack func (msg []byte)
type MsgBroker interface {
	CreateConsumer(fn CallBack) (MsgConsumer, error) // returns a new message producer
	CreateProducer() (MsgProducer, error) // returns a new message consumer
	Close() // closes all resources associated with the brokcer
}

type MsgProducer interface {
	Publish(msg []byte) error
}

type MsgConsumer interface {
	OnMessage(msg []byte)
}

type Closable interface {
	Close()
}

type Msg struct {
	Body      []byte
	Transport interface{}
}