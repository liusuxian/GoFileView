package controller

import (
	v1 "GoFileView/api/v1"
	"GoFileView/internal/code"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"GoFileView/utility/response"
	"GoFileView/utility/utils"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"io/ioutil"
	"strconv"
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
			err = gerror.WrapCode(code.DownloadFailed, err)
			return
		}
		filePath = file
	}
	fileType := gstr.ToLower(gfile.Ext(filePath))

	// MD 文件预览
	if fileType == ".md" {
		dataByte := service.MdPage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// msg 或者 eml 文件预览
	if fileType == ".msg" || fileType == ".eml" {
		var pdfPath string
		pdfPath, err = utils.MsgToPdf(filePath)
		if err != nil {
			err = gerror.WrapCode(code.ConvertToPdfFailed, err)
			return
		}
		var waterPdf string
		waterPdf, err = utils.WaterMark(ctx, pdfPath, req.WaterMark)
		if err != nil {
			err = gerror.WrapCode(code.AddWaterMarkFailed, err)
			return
		}

		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// 后缀是pdf直接读取文件类容返回
	if fileType == ".pdf" {
		var waterPdf string
		waterPdf, err = utils.WaterMark(ctx, filePath, req.WaterMark)
		if err != nil {
			err = gerror.WrapCode(code.AddWaterMarkFailed, err)
			return
		}
		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// 后缀png , jpg ,gif
	if gstr.InArray(service.AllImageEtx, fileType) {
		dataByte := service.ImagePage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// 后缀xlsx
	if (fileType == ".xlsx" || fileType == ".xls") && req.Type != "pdf" {
		dataByte := service.ExcelPage(filePath)
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// 除了PDF外的其他word文件  (如果没有安装ImageMagick，可以将这个分支去掉)
	if gstr.InArray(service.AllOfficeEtx, fileType) && req.Type != "pdf" {
		var pdfPath string
		pdfPath, err = utils.ConvertToPDF(filePath)
		if err != nil {
			err = gerror.WrapCode(code.ConvertToPdfFailed, err)
			return
		}
		var waterPdf string
		waterPdf, err = utils.WaterMark(ctx, pdfPath, req.WaterMark)
		if err != nil {
			err = gerror.WrapCode(code.AddWaterMarkFailed, err)
			return
		}
		var imgPath string
		imgPath, err = utils.ConvertToImg(waterPdf)
		if err != nil {
			err = gerror.WrapCode(code.ConvertToImgFailed, err)
			return
		}
		dataByte := service.OfficePage("cache/convert/" + gfile.Basename(imgPath))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	// 除了PDF外的其他word文件
	if gstr.InArray(service.AllOfficeEtx, fileType) {
		var pdfPath string
		pdfPath, err = utils.ConvertToPDF(filePath)
		if err != nil {
			err = gerror.WrapCode(code.ConvertToPdfFailed, err)
			return
		}
		var waterPdf string
		waterPdf, err = utils.WaterMark(ctx, pdfPath, req.WaterMark)
		if err != nil {
			err = gerror.WrapCode(code.AddWaterMarkFailed, err)
			return
		}
		dataByte := service.PdfPage("cache/pdf/" + gfile.Basename(waterPdf))
		response.HtmlText(ctx, len(dataByte), dataByte)
		return
	}

	err = gerror.WrapCode(code.FileTypeNonsupportPreview, gerror.Newf("%s", fileType))
	return
}

func (c *cView) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if gfile.Exists(req.Url) {
		err = gfile.Remove(req.Url)
		if err != nil {
			return
		}
	}

	allFile, _ := service.GetAllFile("cache/local")
	view := g.RequestFromCtx(ctx).GetView()
	view.Assign("AllFile", allFile)
	err = g.RequestFromCtx(ctx).Response.WriteTpl("resource/template/index/index.html")
	return
}

func (c *cView) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	files := g.RequestFromCtx(ctx).GetUploadFiles("upload-file")
	if files != nil {
		var filenames []string
		filenames, err = files.Save("cache/local")
		if err != nil {
			return
		}

		for _, filename := range filenames {
			oldFilename := "cache/local/" + filename
			newFilename := "cache/local/" + gstr.TrimAll(gfile.Name(filename), "") + gfile.Ext(filename)
			err = gfile.Rename(oldFilename, newFilename)
			if err != nil {
				return
			}
		}
	}

	allFile, _ := service.GetAllFile("cache/local")
	view := g.RequestFromCtx(ctx).GetView()
	view.Assign("AllFile", allFile)
	err = g.RequestFromCtx(ctx).Response.WriteTpl("resource/template/index/index.html")
	return
}

func (c *cView) Img(ctx context.Context, req *v1.ImgReq) (res *v1.ImgRes, err error) {
	logger.Info(ctx, "Img req：", req)
	var dataByte []byte
	dataByte, err = ioutil.ReadFile("cache/download/" + req.Url)
	if err != nil {
		// 如果是本地预览，则文件在local目录下
		dataByte, err = ioutil.ReadFile("cache/local/" + req.Url)
		if err != nil {
			response.HtmlText(ctx, len("404"), []byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
	}

	response.HtmlText(ctx, len(dataByte), dataByte)
	return
}

func (c *cView) Office(ctx context.Context, req *v1.OfficeReq) (res *v1.OfficeRes, err error) {
	logger.Info(ctx, "Office req：", req)
	var dataByte []byte
	dataByte, err = ioutil.ReadFile("cache/convert/" + req.Url)
	if err != nil {
		response.HtmlText(ctx, len("404"), []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return
	}

	response.HtmlText(ctx, len(dataByte), dataByte)
	return
}

func (c *cView) Pdf(ctx context.Context, req *v1.PdfReq) (res *v1.PdfRes, err error) {
	logger.Info(ctx, "Pdf req：", req)
	var dataByte []byte
	dataByte, err = ioutil.ReadFile("cache/pdf/" + req.Url)
	if err != nil {
		response.HtmlText(ctx, len("404"), []byte("出现了一些问题,导致File View无法获取您的数据!"))
		return
	}

	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(dataByte)))
	_, err = g.RequestFromCtx(ctx).Response.Writer.Write(dataByte)
	return
}
