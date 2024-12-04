package svc

import (
	"codeup.aliyun.com/61b84a04fa282c88e1039838/osssdk"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	MsOssModel osssdk.OSSModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MsOssModel: osssdk.NewAliYunOssModel(c.Oss),
	}
}
