package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/delete" tags:"delete" method:"get" summary:"You first delete api"`
	Url    string `json:"Url" dc:"文件url地址"`
}
type DeleteRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
