package config

type CORS struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whiteList" json:"whiteList" yaml:"whiteList"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin" json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `mapstructure:"allowMethods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `mapstructure:"allowHeaders" json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"exposeHeaders" json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allowCredentials" json:"allow-credentials" yaml:"allow-credentials"`
}
