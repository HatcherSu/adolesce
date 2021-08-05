package redis

type Setting func()

type Options struct {
	Network      string `json:"network"` // unix or tcp
	Addr         string `json:"addr"`    // host:port
	Password     string `json:"password"`
	Database     int    `json:"database"`
	PoolSize     int    `json:"pool_size"`
	PoolTimeout  int    `json:"pool_timeout"`
	DialTimeout  int    `json:"dial_timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}
