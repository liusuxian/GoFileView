package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeDownloadFailed            = gcode.New(1000, "文件下载失败", "")       // 文件下载失败
	CodeConvertToPdfFailed        = gcode.New(1001, "转pdf失败", "")       // 转pdf失败
	CodeAddWaterMarkFailed        = gcode.New(1002, "添加水印失败", "")       // 添加水印失败
	CodeConvertToImgFailed        = gcode.New(1003, "转图片失败", "")        // 转图片失败
	CodeFileTypeNonsupportPreview = gcode.New(1004, "暂不支持该类型文件预览！", "") // 暂不支持该类型文件预览！
)
