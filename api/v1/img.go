package v1

import "github.com/gogf/gf/v2/frame/g"

type ImgReq struct {
	g.Meta `path:"/img" tags:"img" method:"get" summary:"You first img api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type ImgRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
