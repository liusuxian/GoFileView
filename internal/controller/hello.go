package controller

import (
	"GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"GoFileView/utility/utils"
	"context"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	logger.Print(ctx, "Hello req：", req)
	resultPath := utils.ConvertToPDF("test.pptx")
	utils.ConvertToImg(resultPath)
	utils.WaterMark(resultPath, "")
	response.JsonOK(ctx, "Hello World!")
	return
}
