package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
	"io/ioutil"
)

var (
	ViewOffice = cViewOffice{}
)

type cViewOffice struct{}

func (c *cViewOffice) Office(ctx context.Context, req *v1.ViewOfficeReq) (res *v1.ViewOfficeRes, err error) {
	logger.Info(ctx, "ViewOffice req：", req)
	imgPath := req.Url
	DataByte, err := ioutil.ReadFile("cache/convert/" + imgPath)
	if err != nil {
		response.HtmlPage(ctx, "404", []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return
	}

	response.HtmlPageOK(ctx, DataByte)
	return
}
