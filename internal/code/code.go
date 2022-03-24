package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	DownloadFailed            = gcode.New(1000, "文件下载失败", "")       // 文件下载失败
	ConvertToPdfFailed        = gcode.New(1001, "转pdf失败", "")       // 转pdf失败
	AddWaterMarkFailed        = gcode.New(1002, "添加水印失败", "")       // 添加水印失败
	ConvertToImgFailed        = gcode.New(1003, "转图片失败", "")        // 转图片失败
	FileTypeNonsupportPreview = gcode.New(1004, "暂不支持该类型文件预览！", "") // 暂不支持该类型文件预览！
)
