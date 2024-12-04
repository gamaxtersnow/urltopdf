package config

import (
	"codeup.aliyun.com/61b84a04fa282c88e1039838/osssdk"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Oss osssdk.OssConf
}
