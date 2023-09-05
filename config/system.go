package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                            // 端口值
	DbType        string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`                      // 数据库类型:mysql(默认)
	OssType       string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"useMultipoint" json:"useMultipoint" yaml:"useMultipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"useRedis" json:"useRedis" yaml:"useRedis"`                // 使用redis
	IPLimitCount  int    `mapstructure:"ipLimitCount" json:"ipLimitCount" yaml:"ipLimitCount"`    // IP限制次数 一个小时15000次
	IPLimitTime   int    `mapstructure:"ipLimitTime" json:"ipLimitTime" yaml:"ipLimitTime"`       // IP限制时间，单位是s
	RouterPrefix  string `mapstructure:"routerPrefix" json:"routerPrefix" yaml:"routerPrefix"`    // 全局路由前缀
}
