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
	var dataByte []byte
	dataByte, err = ioutil.ReadFile("cache/download/" + req.Url)
	if err != nil {
		// 如果是本地预览，则文件在local目录下
		dataByte, err = ioutil.ReadFile("cache/local/" + req.Url)
		if err != nil {
			response.HtmlText(ctx, len("404"), []byte("出现了一些问题,导致File View无法获取您的数据!"))
			return res, nil
		}
	}

	response.HtmlText(ctx, len(dataByte), dataByte)
	return res, nil
}
