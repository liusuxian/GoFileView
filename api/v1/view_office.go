package v1

import "github.com/gogf/gf/v2/frame/g"

type ViewOfficeReq struct {
	g.Meta `path:"/view/office" tags:"office" method:"get" summary:"You first office api"`
	Url    string `json:"Url" dc:"文件url地址"`
}
type ViewOfficeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
