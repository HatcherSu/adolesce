package kafka

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

type AsyncProducerSetter func(*asyncProducer)

type asyncProducer struct {
	mu            sync.Mutex
	log           *zap.Logger
	producer      sarama.AsyncProducer
	topic         string
	errHandle     func(*sarama.ProducerError)
	successHandle func(message *sarama.ProducerMessage)
}

func NewAsyncProducer(o *ProducerOptions, sets ...AsyncProducerSetter) Producer {
	p := &asyncProducer{
		log:   zap.NewExample(),
		topic: o.Topic,
	}
	config := sarama.NewConfig()
	for _, s := range sets {
		s(p)
	}
	// producer set
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Wait for all in-sync replicas to ack the message
	config.Producer.Compression = getCompress(o.Compression) // compression select
	config.Producer.Flush.Frequency = time.Duration(o.FlushFrequency) * time.Millisecond
	if p.successHandle != nil {
		config.Producer.Return.Successes = true
	}

	// net set
	tlsConfig, err := createTlsConfiguration(o.CertFile, o.KeyFile, o.CaFile, o.VerifySsl)
	if err != nil {
		p.log.Error("create tls config error", zap.Error(err))
		return nil
	}
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}

	// create producer
	p.producer, err = sarama.NewAsyncProducer(strings.Split(o.Brokers, ","), config)
	if err != nil {
		p.log.Error("new sync producer error", zap.Error(err))
		return nil
	}

	go func() {
		for {
			select {
			case errMsg := <-p.producer.Errors():
				p.log.Error("produce message error", zap.Any("produce_err", errMsg))
				if p.errHandle != nil {
					p.errHandle(errMsg)
				}
			case msg := <-p.producer.Successes():
				p.log.Info("produce message success", zap.Any("message", msg))
				if p.errHandle != nil {
					p.successHandle(msg)
				}
			}
		}
	}()
	return p
}

// Implementation Producer interface

func (a *asyncProducer) SendMessage(msg Message, header *MessageHeader) error {
	jsonStr, err := message2JsonStr(msg)
	if err != nil {
		a.log.Error("message convert to json error", zap.Error(err))
		return err
	}
	var pMsg sarama.ProducerMessage
	if header != nil {
		pMsg = header.header2Message()
	}
	pMsg.Topic = a.topic
	pMsg.Value = sarama.StringEncoder(jsonStr)
	a.producer.Input() <- &pMsg
	return nil
}

func (a *asyncProducer) SendMessages(msgs []Message, header *MessageHeader) error {
	var pMsg sarama.ProducerMessage
	if header != nil {
		pMsg = header.header2Message()
	}
	pMsg.Topic = a.topic
	for _, msg := range msgs {
		jsonStr, err := message2JsonStr(msg)
		if err != nil {
			a.log.Error("message convert to json error", zap.Error(err))
			return err
		}
		pMsg.Value = sarama.StringEncoder(jsonStr)
		a.producer.Input() <- &pMsg
	}
	return nil
}

func (a *asyncProducer) Topic(topic string) Producer {
	return &asyncProducer{
		log:      a.log,
		producer: a.producer,
		topic:    topic,
	}
}

func (a *asyncProducer) Close() {
	if err := a.producer.Close(); err != nil {
		a.log.Error("async producer close error", zap.Error(err))
	}
}

// WithXxx Structure configuration

func WithAsyncProducerLogger(log *zap.Logger) AsyncProducerSetter {
	return func(p *asyncProducer) {
		p.log = log.Named("[async producer]")
	}
}

func WithErrorsFn(errHandle func(*sarama.ProducerError)) AsyncProducerSetter {
	return func(p *asyncProducer) {
		p.errHandle = errHandle
	}
}

func WithSuccessFn(successHandle func(*sarama.ProducerMessage)) AsyncProducerSetter {
	return func(p *asyncProducer) {
		p.successHandle = successHandle
	}
}
