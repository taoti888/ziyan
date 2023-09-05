package config

type JWT struct {
	SigningKey  string `mapstructure:"signingKey" json:"signingKey" yaml:"signingKey"`    // jwt签名
	ExpiresTime string `mapstructure:"expiresTime" json:"expiresTime" yaml:"expiresTime"` // 过期时间
	BufferTime  string `mapstructure:"bufferTime" json:"bufferTime" yaml:"bufferTime"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                // 签发者
}
