package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewReq struct {
	g.Meta    `path:"/view" tags:"view" method:"get" summary:"You first view api"`
	Url       string `json:"url" dc:"文件url地址"`
	Type      string `json:"type" dc:"判断是图片展示，还是pdf展示"` // 判断是图片展示，还是pdf展示
	FileWay   string `json:"fileWay" dc:"判断是否是本地文件"`    // 判断是否是本地文件
	WaterMark string `json:"waterMark" dc:"水印内容"`       // 水印内容
}
type ViewRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
