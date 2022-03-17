package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Upload = cUpload{}
)

type cUpload struct{}

func (c *cUpload) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	files := g.RequestFromCtx(ctx).GetUploadFile("upload-file")
	if files != nil {
		var filename string
		filename, err = files.Save("cache/local")
		if err != nil {
			logger.Error(ctx, "Upload Save Error:", err.Error())
		}
		oldFilename := "cache/local/" + filename
		newFilename := "cache/local/" + gstr.TrimAll(gfile.Name(filename), "") + gfile.Ext(filename)
		err = gfile.Rename(oldFilename, newFilename)
		if err != nil {
			logger.Error(ctx, "Upload Rename Error:", err.Error())
		}
	}

	allFile, _ := service.GetAllFile("cache/local")
	view := g.RequestFromCtx(ctx).GetView()
	view.Assign("AllFile", allFile)
	err = g.RequestFromCtx(ctx).Response.WriteTpl("resource/template/index/index.html")
	if err != nil {
		logger.Error(ctx, "Upload Error:", err.Error())
	}
	return res, nil
}
