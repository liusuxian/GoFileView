package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewImgReq struct {
	g.Meta `path:"/view/img" tags:"img" method:"get" summary:"You first img api"`
	Url    string `json:"Url" dc:"文件url地址"`
}
type ViewImgRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
