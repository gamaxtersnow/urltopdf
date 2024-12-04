package main

import (
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/config"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/handler"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/svc"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/types"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/urltopdf-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors(), rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.WriteJson(w, http.StatusOK, types.TokenErrorInfo{
			ErrorCode: 401,
			ErrorMsg:  err.Error(),
		})
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (int, any) {
		responseBody := types.ErrorResponse{}
		responseBody.ErrorInfo.ErrorCode = 102
		responseBody.ErrorInfo.ErrorMsg = err.Error()
		return http.StatusOK, responseBody
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
