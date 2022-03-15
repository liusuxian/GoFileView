package v1

import "github.com/gogf/gf/v2/frame/g"

type UploadReq struct {
	g.Meta `path:"/upload" tags:"upload" method:"post" summary:"You first upload api"`
}
type UploadRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
