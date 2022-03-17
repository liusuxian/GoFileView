package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	Delete = cDelete{}
)

type cDelete struct{}

func (c *cDelete) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if gfile.Exists(req.Url) {
		err = gfile.Remove(req.Url)
		if err != nil {
			logger.Error(ctx, "Delete Error: ", err.Error())
		}
	}

	allFile, _ := service.GetAllFile("cache/local")
	view := g.RequestFromCtx(ctx).GetView()
	view.Assign("AllFile", allFile)
	err = g.RequestFromCtx(ctx).Response.WriteTpl("resource/template/index/index.html")
	if err != nil {
		logger.Error(ctx, "Delete Error:", err.Error())
	}

	return res, nil
}
