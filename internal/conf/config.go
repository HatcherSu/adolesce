package conf

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

// Configs 全局配置类
type Configs struct {
	Log      LoggerConfig
	DataBase DatabaseConfig
	Http     HttpServerConfig
	Redis    RedisClientConfig
}

func NewConfig(envPath string) (*Configs, error) {
	var config Configs
	if err := godotenv.Load(envPath); err != nil {
		err = fmt.Errorf("NewConfig->Load: %w", err)
		return nil, err
	}
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// LoggerConfig 全局Logger配置
type LoggerConfig struct {
	Env             string `env:"LOG_ENV"`
	OutputPath      string `env:"LOG_OUTPUT_PATH" yaml:"output_path"`
	ErrorOutputPath string `env:"LOG_ERROR_OUTPUT_PATH" yaml:"error_output_path"`
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
	HttpMode          string `env:"HTTP_MODE" yaml:"mode"`
	HttpPort          int    `env:"HTTP_PORT" yaml:"port"`
	HttpIPAddr        string `env:"HTTP_IP_ADDR" yaml:"ip_addr"`
	DialTimeoutSecond int    `env:"HTTP_DIAL_TIMEOUT_SECOND" yaml:"dial_timeout_second"`
}

// RedisClientConfig redis配置
type RedisClientConfig struct {
	Network      string `env:"REDIS_NETWORK"` // unix or tcp
	Address      string `env:"REDIS_ADDRESS"` // host:port
	Password     string `env:"REDIS_PASSWORD"`
	Database     int    `env:"REDIS_DATABASE"`
	PoolSize     int    `env:"REDIS_POOLSIZE" yaml:"pool_size"`
	PoolTimeout  int    `env:"REDIS_POOLTIMEOUT" yaml:"pool_timeout"`
	DialTimeout  int    `env:"REDIS_DIALTIMEOUT" yaml:"dial_timeout"`
	ReadTimeout  int    `env:"REDIS_READTIMEOUT" yaml:"read_timeout"`
	WriteTimeout int    `env:"REDIS_WRITETIMEOUT" yaml:"write_timeout"`
}
