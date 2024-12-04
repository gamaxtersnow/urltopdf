// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// url转pdf
				Method:  http.MethodPost,
				Path:    "/utils/urlToPdf",
				Handler: UrlToPdfHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}