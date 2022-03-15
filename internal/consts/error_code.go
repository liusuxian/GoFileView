package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeDownloadFailed     = gcode.New(10000, "文件下载失败", "") // 文件下载失败
	CodeConvertToPdfFailed = gcode.New(10001, "转pdf失败", "") // 转pdf失败
	CodeAddWaterMarkFailed = gcode.New(10002, "添加水印失败", "") // "添加水印失败"
	CodeConvertToImgFailed = gcode.New(10003, "转图片失败", "")  // "转图片失败"
)
