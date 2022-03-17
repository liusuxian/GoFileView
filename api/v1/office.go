package v1

import "github.com/gogf/gf/v2/frame/g"

type OfficeReq struct {
	g.Meta `path:"/office" tags:"office" method:"get" summary:"You first office api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type OfficeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
