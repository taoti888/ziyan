package config

type Captcha struct {
	KeyLong            int `mapstructure:"keyLong" json:"keyLong" yaml:"keyLong"`                                  // 验证码长度
	ImgWidth           int `mapstructure:"imgWidth" json:"imgWidth" yaml:"imgWidth"`                               // 验证码宽度
	ImgHeight          int `mapstructure:"imgHeight" json:"imgHeight" yaml:"imgHeight"`                            // 验证码高度
	OpenCaptcha        int `mapstructure:"openCaptcha" json:"openCaptcha" yaml:"openCaptcha"`                      // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"openCaptchaTimeOut" json:"openCaptchaTimeOut" yaml:"openCaptchaTimeOut"` // 防爆破验证码超时时间，单位：s(秒)
}
