package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"io/ioutil"
	"strconv"
)

var (
	Pdf = cPdf{}
)

type cPdf struct{}

func (c *cPdf) Pdf(ctx context.Context, req *v1.PdfReq) (res *v1.PdfRes, err error) {
	logger.Info(ctx, "Pdf req：", req)
	var dataByte []byte
	dataByte, err = ioutil.ReadFile("cache/pdf/" + req.Url)
	if err != nil {
		response.HtmlText(ctx, len("404"), []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return res, nil
	}

	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(dataByte)))
	_, err = g.RequestFromCtx(ctx).Response.Writer.Write(dataByte)
	if err != nil {
		logger.Error(ctx, "Pdf Error:", err.Error())
	}
	return res, nil
}
