package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
	"io/ioutil"
)

var (
	ViewImg = cViewImg{}
)

type cViewImg struct{}

func (c *cViewImg) Img(ctx context.Context, req *v1.ViewImgReq) (res *v1.ViewImgRes, err error) {
	logger.Info(ctx, "ViewImg req：", req)
	imgPath := req.Url
	DataByte, err := ioutil.ReadFile("cache/download/" + imgPath)
	if err != nil {
		// 如果是本地预览，则文件在local目录下
		DataByte, err = ioutil.ReadFile("cache/local/" + imgPath)
		if err != nil {
			response.HtmlPage(ctx, "404", []byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
	}

	response.HtmlPageOK(ctx, DataByte)
	return
}
