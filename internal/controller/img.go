package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
	"io/ioutil"
)

var (
	Img = cImg{}
)

type cImg struct{}

func (c *cImg) Img(ctx context.Context, req *v1.ImgReq) (res *v1.ImgRes, err error) {
	logger.Info(ctx, "Img req：", req)
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