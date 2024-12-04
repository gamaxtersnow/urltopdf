package handler

import (
	"net/http"

	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/logic"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/svc"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// UrlToPdfHandler url转pdf
func UrlToPdfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UrlToPdfRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUrlToPdfLogic(r.Context(), svcCtx)
		resp, err := l.UrlToPdf(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
