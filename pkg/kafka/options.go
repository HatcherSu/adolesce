package kafka

import (
	"github.com/spf13/pflag"
)

type ConsumerOptions struct {
	Brokers  string `json:"brokers"` //  multiple brokers are separated by ’,‘
	Group    string `json:"group"`
	Version  string `json:"version"`
	Topics   string `json:"topics"` // multiple topic are separated by ’,‘
	Assignor string `json:"assignor"`
	Oldest   bool   `json:"oldest"`
	Verbose  bool   `json:"verbose"`
}

func (o *ConsumerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	fs.StringVar(&o.Group, "group", "", "Kafka consumer group definition")
	fs.StringVar(&o.Version, "version", "2.1.1", "Kafka cluster version")
	fs.StringVar(&o.Topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
	fs.StringVar(&o.Assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	fs.BoolVar(&o.Oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	fs.BoolVar(&o.Verbose, "verbose", false, "Sarama logging")
}

type ProducerOptions struct {
	Brokers        string `json:"brokers"` //  multiple brokers are separated by ’,‘
	Verbose        bool   `json:"verbose"`
	Topic          string `json:"topic"`
	RetryMax       int    `json:"retry_max"`
	CertFile       string `json:"cert_file"`
	KeyFile        string `json:"key_file"`
	CaFile         string `json:"ca_file"`
	VerifySsl      bool   `json:"verify_ssl"`
	FlushFrequency int    `json:"flush_frequency"` // eg: 500ms
	Compression    string `json:"compression"`     // none,gzip,snappy,lz4,zstd
	Successes      bool   `json:"successes"`
}

