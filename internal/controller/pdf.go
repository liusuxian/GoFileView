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
	DataByte, err := ioutil.ReadFile("cache/pdf/" + req.Url)
	if err != nil {
		response.HtmlPage(ctx, "404", []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return
	}

	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(DataByte)))
	_, err = g.RequestFromCtx(ctx).Response.Writer.Write(DataByte)
	if err != nil {
		logger.Error(ctx, "Pdf Error:", err.Error())
	}
	return
}
