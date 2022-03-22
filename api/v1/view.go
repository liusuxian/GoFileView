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

type UploadReq struct {
	g.Meta `path:"/upload" tags:"upload" method:"post" summary:"You first upload api"`
}
type UploadRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type DeleteReq struct {
	g.Meta `path:"/delete" tags:"delete" method:"get" summary:"You first delete api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type DeleteRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ImgReq struct {
	g.Meta `path:"/img" tags:"img" method:"get" summary:"You first img api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type ImgRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type OfficeReq struct {
	g.Meta `path:"/office" tags:"office" method:"get" summary:"You first office api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type OfficeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PdfReq struct {
	g.Meta `path:"/pdf" tags:"pdf" method:"get" summary:"You first pdf api"`
	Url    string `json:"url" dc:"文件url地址"`
}
type PdfRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
