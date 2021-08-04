package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type Message map[string]interface{}
type MessageHeader struct {
	Partition int32
	Offset    int64
	Headers   map[string]string
	Key       string
	Metadata  interface{}
	Timestamp time.Time
}

type MsgCodeType string

//  todo
const (
	XML  MsgCodeType = "xml"
	JSON MsgCodeType = "json"
)

type Producer interface {
	SendMessage(msg Message, header *MessageHeader) error
	SendMessages(msgs []Message, header *MessageHeader) error
	Topic(string) Producer
	Close()
}

type SyncProducerSetter func(*syncProducer)

type syncProducer struct {
	mu        sync.Mutex
	log       *zap.Logger
	producer  sarama.SyncProducer
	topic     string
	Partition int32 // set after send message
	Offset    int64 // set after send message
}

func NewSyncProducer(o *ProducerOptions, sets ...SyncProducerSetter) Producer {
	p := &syncProducer{
		log:   zap.NewExample(),
		topic: o.Topic,
	}
	config := sarama.NewConfig()
	for _, s := range sets {
		s(p)
	}
	// Wait for all in-sync replicas to ack the message
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = getCompress(o.Compression)
	config.Producer.Flush.Frequency = time.Duration(o.FlushFrequency) * time.Millisecond
	config.Producer.Retry.Max = o.RetryMax
	config.Producer.Return.Successes = true
	tlsConfig, err := createTlsConfiguration(o.CertFile, o.KeyFile, o.CaFile, o.VerifySsl)
	if err != nil {
		p.log.Error("create tls config error", zap.Error(err))
		return nil
	}
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}
	producer, err := sarama.NewSyncProducer(strings.Split(o.Brokers, ","), config)
	if err != nil {
		p.log.Error("new sync producer error", zap.Error(err))
		return nil
	}
	p.producer = producer
	return p
}

// Implementation Producer interface

func (s *syncProducer) SendMessage(message Message, header *MessageHeader) error {
	jsonStr, err := message2JsonStr(message)
	if err != nil {
		s.log.Error("message convert to json error", zap.Error(err))
		return err
	}
	var msg sarama.ProducerMessage
	if header != nil {
		msg = header.header2Message()
	}
	msg.Topic = s.topic
	msg.Value = sarama.StringEncoder(jsonStr)
	partition, offset, err := s.producer.SendMessage(&msg)
	if err != nil {
		return fmt.Errorf("producer send message err, %v", err)
	}
	s.Partition = partition
	s.Offset = offset
	return nil
}

func (s *syncProducer) SendMessages(messages []Message, header *MessageHeader) error {
	var list []*sarama.ProducerMessage
	var pMsg sarama.ProducerMessage
	if header != nil {
		pMsg = header.header2Message()
	}
	pMsg.Topic = s.topic
	for _, msg := range messages {
		jsonStr, err := message2JsonStr(msg)
		if err != nil {
			s.log.Error("message convert to json error", zap.Error(err))
			return err
		}
		pMsg.Value = sarama.StringEncoder(jsonStr)
		list = append(list, &pMsg)
	}
	return s.producer.SendMessages(list)
}

func (s *syncProducer) Close() {
	if err := s.producer.Close(); err != nil {
		s.log.Error("sync producer close error", zap.Error(err))
	}
}

func (s *syncProducer) Topic(topic string) Producer {
	return &syncProducer{
		log:      s.log,
		producer: s.producer,
		topic:    topic,
	}
}

// WithXxx Structure configuration

func WithProducerLogger(log *zap.Logger) SyncProducerSetter {
	return func(p *syncProducer) {
		p.log = log.Named("[sync producer]")
	}
}

// other method

// create tls config with file
func createTlsConfiguration(certFile, keyFile, caFile string, verifySsl bool) (t *tls.Config, err error) {
	if certFile == "" || keyFile != "" || caFile != "" {
		return nil, nil
	}
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("LoadX509KeyPair err %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("ReadFile err,CaFile: %s,err: %v", caFile, err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	t = &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: verifySsl,
	}
	// will be nil by default if nothing is provided
	return t, nil
}

func message2JsonStr(msg Message) (string, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("message to json error:%v", err)
	}
	return string(b), nil
}

var compressMap = map[string]sarama.CompressionCodec{
	"none":   sarama.CompressionNone,
	"gzip":   sarama.CompressionGZIP,
	"snappy": sarama.CompressionSnappy,
	"lz4":    sarama.CompressionLZ4,
	"zstd":   sarama.CompressionZSTD,
}

func getCompress(compress string) sarama.CompressionCodec {
	v, ok := compressMap[compress]
	if ok {
		return v
	}
	return sarama.CompressionNone
}

func (m MessageHeader) header2Message() sarama.ProducerMessage {
	var headers []sarama.RecordHeader

	for k, v := range m.Headers {
		headers = append(headers, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}
	return sarama.ProducerMessage{
		Key:       sarama.StringEncoder(m.Key),
		Headers:   headers,
		Metadata:  m.Metadata,
		Offset:    m.Offset,
		Partition: m.Partition,
		Timestamp: m.Timestamp,
	}
}
