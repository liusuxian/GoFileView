package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
	"io/ioutil"
)

var (
	ViewPdf = cViewPdf{}
)

type cViewPdf struct{}

func (c *cViewPdf) Pdf(ctx context.Context, req *v1.ViewPdfReq) (res *v1.ViewPdfRes, err error) {
	logger.Info(ctx, "ViewPdf req：", req)
	imgPath := req.Url
	DataByte, err := ioutil.ReadFile("cache/pdf/" + imgPath)
	if err != nil {
		response.HtmlPage(ctx, "404", []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return
	}

	response.HtmlPageOK(ctx, DataByte)
	return
}
