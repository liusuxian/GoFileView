package controller

import (
	"GoFileView/api/v1"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"context"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	logger.Info(ctx, "Hello req：", req)
	response.JsonOK(ctx, "Hello World!")
	return
}
