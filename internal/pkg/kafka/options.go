package kafka

type Options struct {
	Brokers  string // Kafka bootstrap brokers to connect to, as a comma separated list
	group    string
	version  string
	topics   string
	assignor string
	oldest   bool
	verbose  bool
}
