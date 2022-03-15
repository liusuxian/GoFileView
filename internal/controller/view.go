package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/consts"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"GoFileView/utility/utils"
	"context"
	"path"
	"strings"
)

// 本地文件路径
var filePath string

var (
	View = cView{}
)

type cView struct{}

func (c *cView) View(ctx context.Context, req *v1.ViewReq) (res *v1.ViewRes, err error) {
	logger.Info(ctx, "View req：", req)
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
	// MD 文件预览
	if fileType == ".md" {
		dataByte := service.MdPage(filePath)
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// msg 或者 eml 文件预览
	if fileType == ".msg" || fileType == ".eml" {
		pdfPath := utils.MsgToPdf(filePath)
		if pdfPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToPdfFailed.Code(), consts.CodeConvertToPdfFailed.Message())
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.JsonExit(ctx, consts.CodeAddWaterMarkFailed.Code(), consts.CodeAddWaterMarkFailed.Message())
		}

		dataByte := service.PdfPage("cache/pdf/" + path.Base(waterPdf))
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// 后缀是pdf直接读取文件类容返回
	if fileType == ".pdf" {
		waterPdf := utils.WaterMark(filePath, req.WaterMark)
		if waterPdf == "" {
			response.JsonExit(ctx, consts.CodeAddWaterMarkFailed.Code(), consts.CodeAddWaterMarkFailed.Message())
		}
		logger.Info(ctx, "waterPdf: ", waterPdf)
		dataByte := service.PdfPage("cache/pdf/" + path.Base(waterPdf))
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// 后缀png , jpg ,gif
	if utils.IsInArr(fileType, service.AllImageEtx) {
		dataByte := service.ImagePage(filePath)
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// 后缀xlsx
	if (fileType == ".xlsx" || fileType == ".xls") && req.Type != "pdf" {
		dataByte := service.ExcelPage(filePath)
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// 除了PDF外的其他word文件  (如果没有安装ImageMagick，可以将这个分支去掉)
	if utils.IsInArr(fileType, service.AllOfficeEtx) && req.Type != "pdf" {
		pdfPath := utils.ConvertToPDF(filePath)
		logger.Debug(ctx, "pdfPath: ", pdfPath)
		if pdfPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToPdfFailed.Code(), consts.CodeConvertToPdfFailed.Message())
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		logger.Debug(ctx, "waterPdf: ", waterPdf)
		if waterPdf == "" {
			response.JsonExit(ctx, consts.CodeAddWaterMarkFailed.Code(), consts.CodeAddWaterMarkFailed.Message())
		}
		imgPath := utils.ConvertToImg(waterPdf)
		logger.Debug(ctx, "imgPath: ", imgPath)
		if imgPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToImgFailed.Code(), consts.CodeConvertToImgFailed.Message())
		}
		dataByte := service.OfficePage("cache/convert/" + path.Base(imgPath))
		response.HtmlPageOK(ctx, dataByte)
		return
	}
	// 除了PDF外的其他word文件
	if utils.IsInArr(fileType, service.AllOfficeEtx) {
		pdfPath := utils.ConvertToPDF(filePath)
		if pdfPath == "" {
			response.JsonExit(ctx, consts.CodeConvertToPdfFailed.Code(), consts.CodeConvertToPdfFailed.Message())
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.JsonExit(ctx, consts.CodeAddWaterMarkFailed.Code(), consts.CodeAddWaterMarkFailed.Message())
		}
		dataByte := service.PdfPage("cache/pdf/" + path.Base(waterPdf))
		response.HtmlPageOK(ctx, dataByte)
		return
	}

	response.JsonExit(ctx, consts.CodeFileTypeNonsupportPreview.Code(), consts.CodeFileTypeNonsupportPreview.Message())
	return
}
