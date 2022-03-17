package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/consts"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"GoFileView/utility/utils"
	"context"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
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
		var file string
		file, err = utils.DownloadFile(req.Url, "cache/download/"+gfile.Basename(req.Url))
		if err != nil {
			logger.Info(ctx, "下载文件出错: ", err.Error())
			response.Json(ctx, consts.CodeDownloadFailed, nil)
			return res, nil
		}
		filePath = file
	}
	fileType := gstr.ToLower(gfile.Ext(filePath))

	// MD 文件预览
	if fileType == ".md" {
		dataByte := service.MdPage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// msg 或者 eml 文件预览
	if fileType == ".msg" || fileType == ".eml" {
		pdfPath := utils.MsgToPdf(filePath)
		if pdfPath == "" {
			response.Json(ctx, consts.CodeConvertToPdfFailed, nil)
			return res, nil
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.Json(ctx, consts.CodeAddWaterMarkFailed, nil)
			return res, nil
		}

		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// 后缀是pdf直接读取文件类容返回
	if fileType == ".pdf" {
		waterPdf := utils.WaterMark(filePath, req.WaterMark)
		if waterPdf == "" {
			response.Json(ctx, consts.CodeAddWaterMarkFailed, nil)
			return res, nil
		}
		logger.Info(ctx, "waterPdf: ", waterPdf)
		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// 后缀png , jpg ,gif
	if gstr.InArray(service.AllImageEtx, fileType) {
		dataByte := service.ImagePage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// 后缀xlsx
	if (fileType == ".xlsx" || fileType == ".xls") && req.Type != "pdf" {
		dataByte := service.ExcelPage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// 除了PDF外的其他word文件  (如果没有安装ImageMagick，可以将这个分支去掉)
	if gstr.InArray(service.AllOfficeEtx, fileType) && req.Type != "pdf" {
		pdfPath := utils.ConvertToPDF(filePath)
		if pdfPath == "" {
			response.Json(ctx, consts.CodeConvertToPdfFailed, nil)
			return res, nil
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.Json(ctx, consts.CodeAddWaterMarkFailed, nil)
			return res, nil
		}
		imgPath := utils.ConvertToImg(waterPdf)
		if imgPath == "" {
			response.Json(ctx, consts.CodeConvertToImgFailed, nil)
			return res, nil
		}
		dataByte := service.OfficePage("cache/convert/" + gfile.Basename(imgPath))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	// 除了PDF外的其他word文件
	if gstr.InArray(service.AllOfficeEtx, fileType) {
		pdfPath := utils.ConvertToPDF(filePath)
		if pdfPath == "" {
			response.Json(ctx, consts.CodeConvertToPdfFailed, nil)
			return res, nil
		}
		waterPdf := utils.WaterMark(pdfPath, req.WaterMark)
		if waterPdf == "" {
			response.Json(ctx, consts.CodeAddWaterMarkFailed, nil)
		}
		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return res, nil
	}

	response.Json(ctx, consts.CodeFileTypeNonsupportPreview, nil)
	return res, nil
}
