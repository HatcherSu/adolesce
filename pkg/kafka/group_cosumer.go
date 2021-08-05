package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"strings"
	"sync"
)

type MessageHandler func(*sarama.ConsumerMessage) error

type ConsumerSetter func(*groupConsumer)

type Consumer interface {
	StartConsume(context.Context)
	Close()
}

type groupConsumer struct {
	topics    string
	mu        sync.Mutex
	wg        *sync.WaitGroup
	handleMsg MessageHandler
	log       *zap.Logger
	client    sarama.ConsumerGroup
	// todo number of consumers
	ConsumeNum int
}

func NewGroupConsumer(o *ConsumerOptions, os ...ConsumerSetter) Consumer {
	c := &groupConsumer{
		log: zap.NewExample(),
	}
	for _, o := range os {
		o(c)
	}
	c.topics = o.Topics
	if o.Verbose && c.log != nil {
		sarama.Logger = zap.NewStdLog(c.log)
	}
	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(o.Version)
	if err != nil {
		c.log.Error("Error parsing Kafka version: %v", zap.Error(err))
		return nil
	}
	config.Version = version

	switch o.Assignor {
	case sarama.StickyBalanceStrategyName:
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case sarama.RoundRobinBalanceStrategyName:
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case sarama.RangeBalanceStrategyName:
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		c.log.Error("Unrecognized consumer group partition assignor", zap.String("Assignor", o.Assignor))
		return nil
	}

	if o.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	client, err := sarama.NewConsumerGroup(strings.Split(o.Brokers, ","), o.Group, config)
	if err != nil {
		c.log.Error("Error creating consumer group client", zap.Error(err))
		return nil
	}
	c.client = client
	return c
}

func WithConsumerLogger(log *zap.Logger) ConsumerSetter {
	return func(g *groupConsumer) {
		g.log = log.Named("[consumer]")
	}
}

func (c *groupConsumer) StartConsume(ctx context.Context) {
	c.mu.Lock()
	defer c.mu.Unlock()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.wg = wg
	go func() {
		defer wg.Done()
		for {
			if err := c.client.Consume(ctx, strings.Split(c.topics, ","), c); err != nil {
				c.log.Error("Error from Consumer", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
}

func (c *groupConsumer) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.wg.Wait() // wait for consume finish
	if err := c.client.Close(); err != nil {
		c.log.Error("consumer close error", zap.Error(err))
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *groupConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *groupConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *groupConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := c.handleMsg(msg); err != nil {
			return err
		}
		// Mark message is consume
		session.MarkMessage(msg, "")
	}
	return nil
}
