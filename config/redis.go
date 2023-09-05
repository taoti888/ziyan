package config

import "time"

type Redis struct {
	Addrs        []string      `mapstructure:"addrs" json:"addrs" yaml:"addrs"`                      // 服务器地址:端口列表,集群模式
	Username     string        `mapstructure:"username" json:"username" yaml:"username"`             // 用户名
	Password     string        `mapstructure:"password" json:"password" yaml:"password"`             // 密码
	PoolSize     int           `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`             // 连接池，单个节点，不是整个集群,默认5倍CPU数
	MaxRedirects int           `mapstructure:"maxRedirects" json:"maxRedirect" yaml:"maxRedirect"`   // 当遇到网络错误或者MOVED/ASK重定向命令时，最多重试几次，默认8
	MinIdleConns int           `mapstructure:"minIdleConns" json:"MinIdleConns" yaml:"minIdleConns"` // 最小闲置连接数，
	IdleTimeout  time.Duration `mapstructure:"idleTimeout" json:"idleTimeout" yaml:"idleTimeout"`    // 空闲连接超时时间
	DialTimeout  time.Duration `mapstructure:"dialTimeout" json:"dialTimeout" yaml:"dialTimeout"`    // 连接超时时间
	ReadTimeout  time.Duration `mapstructure:"readTimeout" json:"readTimeout" yaml:"readTimeout"`    // 读取超时时间
	WriteTimeout time.Duration `mapstructure:"writeTimeout" json:"writeTimeout" yaml:"writeTimeout"` // 写入超时时间
}
