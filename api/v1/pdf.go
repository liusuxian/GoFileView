package v1

import "github.com/gogf/gf/v2/frame/g"

type PdfReq struct {
	g.Meta `path:"/pdf" tags:"pdf" method:"get" summary:"You first pdf api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type PdfRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
