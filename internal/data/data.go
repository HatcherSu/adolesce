package data

import (
	"cloud_callback/internal/conf"
	"cloud_callback/internal/pkg/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCallbackLogRepo, NewCallbackInfoRepo)

const connectFormat = "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t"

var db *gorm.DB

type Data struct {
}

func NewData(config *conf.Configs, logger log.Logger) (*Data, func(), error) {
	data := &Data{}
	if err := initDB(config, logger); err != nil {
		return nil, nil, err
	}
	return data, func() {
		_ = data.Close()
	}, nil
}

func initDB(config *conf.Configs, logger log.Logger) (err error) {
	dns := fmt.Sprintf(connectFormat,
		config.DataConf.UserName,
		config.DataConf.Password,
		config.DataConf.Host,
		config.DataConf.Database,
		config.DataConf.Charset,
		config.DataConf.ParseTime)

	db, err = gorm.Open("mysql", dns)
	if err != nil {
		logger.Error("initDB-->Open", zap.String("dns", dns), zap.Error(err))
		return err
	}
	db.SetLogger(log.StdInfoLogger())
	db.LogMode(true)
	db.SingularTable(true)
	return nil
}

func (*Data) GetDB() *gorm.DB {
	return db
}

func (*Data) Close() error {
	// todo log
	return db.Close()
}
