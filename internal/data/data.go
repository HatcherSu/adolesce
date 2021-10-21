package data

import (
	"adolesce/internal/conf"
	"adolesce/pkg/log"
	"fmt"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCallbackLogRepo, NewCallbackInfoRepo)

const connectFormat = "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t"

type Data struct {
	db     *gorm.DB
	logger log.Logger
}

func NewData(config *conf.Configs, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		logger: logger,
	}
	db, err := initDB(config, logger)
	if err != nil {
		return nil, nil, err
	}
	data.db = db
	return data, func() {
		_ = data.Close()
	}, nil
}

func initDB(config *conf.Configs, logger log.Logger) (*gorm.DB, error) {
	dns := fmt.Sprintf(connectFormat,
		config.DataBase.UserName,
		config.DataBase.Password,
		config.DataBase.Host,
		config.DataBase.Database,
		config.DataBase.Charset,
		config.DataBase.ParseTime)

	db, err := gorm.Open("mysql", dns)
	if err != nil {
		logger.Error("initDB-->Open", zap.String("dns", dns), zap.Error(err))
		return nil, err
	}
	db.SetLogger(log.StdInfoLogger())
	db.LogMode(true)
	db.SingularTable(true)
	return db, nil
}


func initRedis(config *conf.Configs) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:      config.Redis.Network,
		Addr:         config.Redis.Address,
		DB:           config.Redis.Database,
		PoolSize:     config.Redis.PoolSize,
		Password:     config.Redis.Password,
		PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
		DialTimeout:  time.Duration(config.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Redis.WriteTimeout) * time.Second,
	})
	rdb.AddHook(redisotel.TracingHook{})
	return rdb
}


func (d *Data) Close() (err error) {
	if err = d.db.Close(); err != nil {
		d.logger.Error("db-->Close", zap.Error(err))
	}
	return
}
