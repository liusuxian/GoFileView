package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/consts"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"GoFileView/utility/utils"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"path"
	"strconv"
	"strings"
)

// 本地文件路径
var filePath string

var (
	View = cView{}
)

type cView struct{}

func (c *cView) View(ctx context.Context, req *v1.ViewViewReq) (res *v1.ViewViewRes, err error) {
	logger.Print(ctx, "View req：", req)
	if req.FileWay == "local" {
		// 本地文件预览
		filePath = req.Url
	} else {
		// 下载文件
		file, err := utils.DownloadFile(req.Url, "cache/download/"+path.Base(req.Url))
		if err != nil {
			logger.Info(ctx, "下载文件出错: ", err.Error())
			response.JsonExit(ctx, consts.CodeDownloadFailed.Code(), consts.CodeDownloadFailed.Message())
		}
		filePath = file
	}
	fileType := strings.ToLower(path.Ext(filePath))
	// 除了PDF外的其他word文件  (如果没有安装ImageMagick，可以将这个分支去掉)
	if utils.IsInArr(fileType, service.AllOfficeEtx) && req.Type != "pdf" {
		pdfPath := utils.ConvertToPDF(filePath)
		if pdfPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToPdfFailed.Code(), consts.CodeConvertToPdfFailed.Message())
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.JsonExit(ctx, consts.CodeAddWaterMarkFailed.Code(), consts.CodeAddWaterMarkFailed.Message())
		}
		imgPath := utils.ConvertToImg(waterPdf)
		if imgPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToImgFailed.Code(), consts.CodeConvertToImgFailed.Message())
		}
		dataByte := service.OfficePage("cache/convert/" + path.Base(imgPath))
		logger.Print(ctx, dataByte)
		g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(dataByte)))
		g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-type:", "text/html;charset=UTF-8")
		g.RequestFromCtx(ctx).Response.Writer.Write(dataByte)
		return
	}

	response.JsonExit(ctx, consts.CodeFileTypeNonsupportPreview.Code(), consts.CodeFileTypeNonsupportPreview.Message())
	return
}
