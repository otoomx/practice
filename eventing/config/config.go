//Package config provides a basic configuration structure for project AMQP broker settings
package config

type Config struct {
	Broker             string
	Exchange           string
	ExchangeType       string
	BindingKey         string
	PrefetchCount      int
	MaxRetryCount      int
	MsgTtl             int
	Persistent         bool
	UnProcessableQueue string
}


var (
	//default configuration
	//TODO:provide external file or enviroment based config
	cfg = &Config{
		"amqp://guest:guest@localhost:5672/",
		"events.exchange",
		"direct",
		"event",
		2,
		100,
		100000,
		true,
		"unprocessable.queue",
	}
)

// Get returns application AMQP configuration
func Get() *Config {
	return cfg
}
