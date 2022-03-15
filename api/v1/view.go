package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewViewReq struct {
	g.Meta    `path:"/view/view" tags:"view" method:"post" summary:"You first view api"`
	Url       string `json:"Url" dc:"文件url地址"`
	Type      string `json:"Type" v:"pdf、\"\"" dc:"判断是图片展示，还是pdf展示"` // 判断是图片展示，还是pdf展示
	FileWay   string `json:"FileWay" v:"local、\"\"" dc:"判断是否是本地文件"`  // 判断是否是本地文件
	WaterMark string `json:"watermark" dc:"水印内容"`                    // 水印内容
}
type ViewViewRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
