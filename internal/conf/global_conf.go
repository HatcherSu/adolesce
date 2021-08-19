package conf

import (
	"fmt"
	"os"
)

// Configs 全局配置类
type Configs struct {
	LogConf   LoggerConfig
	DataConf  DatabaseConfig
	HttpConf  HttpServerConfig
	RedisConf RedisClientConfig
}

func NewConfig(confObj *Conf) (*Configs, error) {
	if err := confObj.Load(); err != nil {
		err = fmt.Errorf("NewConfig->Load: %w", err)
		return nil, err
	}
	defer func() {
		if err := confObj.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "NewConfig->confObj.Close: %v\n", err)
		}
	}()

	confObj.SetWatchErrHandleFunc(func(err error) {
		_, _ = fmt.Fprintf(os.Stderr, "Conf.WatchErrFunc: %v\n", err)
	})

	if err := confObj.Watch(); err != nil {
		err = fmt.Errorf("NewConfig->Watch: %w", err)
		return nil, err
	}

	// 获取Logger配置
	var loggerConfig LoggerConfig
	if err := confObj.Scan(&loggerConfig, "env"); err != nil {
		err = fmt.Errorf("confObj->Scan->LoggerConfig: %w", err)
		return nil, err
	}
	// 获取数据库配置
	var dataConfig DatabaseConfig
	if err := confObj.Scan(&dataConfig, "env"); err != nil {
		err = fmt.Errorf("confObj->Scan->DatabaseConfig: %w", err)
		return nil, err
	}

	// 获取数据库配置
	var redisConf RedisClientConfig
	if err := confObj.Scan(&redisConf, "env"); err != nil {
		err = fmt.Errorf("confObj->Scan->RedisClientConfig: %w", err)
		return nil, err
	}
	// 获取http配置
	var httpConfig HttpServerConfig
	if err := confObj.Scan(&httpConfig, "env"); err != nil {
		err = fmt.Errorf("confObj->Scan->HttpServerConfig: %w", err)
		return nil, err
	}
	return &Configs{
		loggerConfig,
		dataConfig,
		httpConfig,
		redisConf,
	}, nil
}

// LoggerConfig 全局Logger配置
type LoggerConfig struct {
	Level             string `env:"LOG_LEVEL"`
	Development       bool   `env:"LOG_DEVELOPMENT"`
	Format            string `env:"LOG_FORMAT"`
	EnableColor       bool   `env:"LOG_ENABLE_COLOR"`
	DisableCaller     bool   `env:"LOG_DISABLE_CALLER"`
	DisableStacktrace bool   `env:"LOG_DISABLE_STACKTRACE"`
	OutputPath        string `env:"LOG_OUTPUT_PATH"`
	ErrorOutputPath   string `env:"LOG_ERROR_OUTPUT_PATH"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host      string `env:"DB_HOST"`
	Port      string `env:"DB_PORT"`
	UserName  string `env:"DB_USERNAME"`
	Password  string `env:"DB_PASSWORD"`
	Database  string `env:"DB_DATABASE"`
	ParseTime bool   `env:"DB_PARSE_TIME"`
	Charset   string `env:"DB_CHARSET"`
}

// HttpServerConfig HTTP配置
type HttpServerConfig struct {
	HttpMode          string `env:"HTTP_MODE"`
	HttpPort          int    `env:"HTTP_PORT"`
	HttpIPAddr        string `env:"HTTP_IP_ADDR"`
	DialTimeoutSecond int    `env:"HTTP_DIAL_TIMEOUT_SECOND"`
}

// RedisClientConfig redis配置
type RedisClientConfig struct {
	Network      string `env:"REDIS_NETWORK"` // unix or tcp
	Address      string `env:"REDIS_ADDRESS"` // host:port
	Password     string `env:"REDIS_PASSWORD"`
	Database     int    `env:"REDIS_DATABASE"`
	PoolSize     int    `env:"REDIS_POOLSIZE"`
	PoolTimeout  int    `env:"REDIS_POOLTIMEOUT"`
	DialTimeout  int    `env:"REDIS_DIALTIMEOUT"`
	ReadTimeout  int    `env:"REDIS_READTIMEOUT"`
	WriteTimeout int    `env:"REDIS_WRITETIMEOUT"`
}
