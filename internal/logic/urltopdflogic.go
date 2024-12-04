package logic

import (
	"codeup.aliyun.com/61b84a04fa282c88e1039838/utils"
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"strings"
	"time"

	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/svc"
	"codeup.aliyun.com/61b84a04fa282c88e1039838/urltopdf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UrlToPdfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUrlToPdfLogic url转pdf
func NewUrlToPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UrlToPdfLogic {
	return &UrlToPdfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UrlToPdfLogic) UrlToPdf(req *types.UrlToPdfRequest) (resp *types.UrlToPdfResponse, err error) {
	ctx, cancel := chromedp.NewContext(l.ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var buf []byte
	err = chromedp.Run(ctx,
		// 启用网络拦截
		network.Enable(),
		// 设置请求头
		chromedp.ActionFunc(func(ctx context.Context) error {
			headers := network.Headers{}
			if req.Token != "" {
				headers[req.HeaderTokenName] = req.Token
			}
			return network.SetExtraHTTPHeaders(headers).Do(ctx)
		}),
		chromedp.Navigate(req.Url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(time.Duration(req.Delay)*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().WithPrintBackground(true).
				WithScale(req.Scale).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return nil, err
	}
	objectKey := strings.TrimLeft(utils.RemoveAllExtensions(req.ObjectKey), "/") + ".pdf"
	err = l.svcCtx.MsOssModel.UploadOssByFile(l.ctx, objectKey, buf)
	if err != nil {
		return nil, err
	}
	urlToPdfResponse := &types.UrlToPdfResponse{}
	urlToPdfResponse.ErrorInfo.ErrorCode = 1
	urlToPdfResponse.ErrorInfo.ErrorMsg = "生成成功"
	urlToPdfResponse.Data.FileUrl = strings.Replace(l.svcCtx.Config.Oss.OssEndpoint, "://", "://"+l.svcCtx.Config.Oss.BucketName+".", 1) + "/" + objectKey
	if req.IsSign {
		urlToPdfResponse.Data.FileUrl, _ = l.svcCtx.MsOssModel.GetSignUrl(l.ctx, objectKey, req.SignExpire)
	}
	urlToPdfResponse.Data.ObjecktId = objectKey
	return urlToPdfResponse, nil
}
