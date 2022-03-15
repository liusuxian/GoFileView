package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewReq struct {
	g.Meta    `path:"/view" tags:"view" method:"post" summary:"You first view api"`
	Url       string `json:"Url"`
	Type      string `json:"Type"`      // 判断是图片展示，还是pdf展示
	FileWay   string `json:"FileWay"`   // 判断是否是本地文件
	WaterMark string `json:"watermark"` // 水印
}
type ViewRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
