package config

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`                      // 连接Endpoint
	AccessKeyId     string `mapstructure:"accessKeyID" json:"accessKeyID" yaml:"accessKeyID"`             // 阿里云密钥ID
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret" yaml:"accessKeySecret"` // 阿里云密钥Secret
	BucketName      string `mapstructure:"bucketName" json:"bucketName" yaml:"bucketName"`                // Bucket名称
	BucketUrl       string `mapstructure:"bucketUrl" json:"bucketUrl" yaml:"bucketUrl"`                   // Bucket的url地址
	BasePath        string `mapstructure:"basePath" json:"basePath" yaml:"basePath"`                      // 文件路径
}
