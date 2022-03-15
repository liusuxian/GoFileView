package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewPdfReq struct {
	g.Meta `path:"/view/pdf" tags:"pdf" method:"get" summary:"You first pdf api"`
	Url    string `json:"Url" dc:"文件url地址"`
}
type ViewPdfRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
