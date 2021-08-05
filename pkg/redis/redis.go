package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	*redis.Client
	ctx       context.Context
	defaultDB int
}

func NewRedisClient(o *Options) *Client {
	rdb := redis.NewClient(&redis.Options{
		Network:      o.Network,
		Addr:         o.Addr,
		DB:           o.Database,
		PoolSize:     o.PoolSize,
		PoolTimeout:  time.Duration(o.PoolTimeout) * time.Second,
		DialTimeout:  time.Duration(o.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(o.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(o.WriteTimeout) * time.Second,
	})
	c := &Client{
		Client:    rdb,
		defaultDB: o.Database,
		ctx:       context.Background(),
	}
	return c
}

func (c *Client) Select(index int) *Client {
	c.Pipeline().Select(c.ctx, index)
	return c
}
