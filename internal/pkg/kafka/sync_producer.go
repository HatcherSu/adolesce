package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"sync"
)

type Message map[string]interface{}

type Producer interface {
	SendMessage(Message)
	SendMessageList([]Message)
	Close()
}

type syncProducer struct {
	mu  sync.Mutex
	log *zap.Logger
}

// todo
//func NewSyncProducer(o *ProducerOptions) Producer {
//	sProd := &syncProducer{}
//	config := sarama.NewConfig()
//	// Wait for all in-sync replicas to ack the message
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	config.Producer.Retry.Max = o.RetryMax
//	config.Producer.Return.Successes = true
//	tlsConfig, err := createTlsConfiguration()
//
//}

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
