package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
	"path"
	"strings"
)

// WaterMark pdf增加水印
func WaterMark(pdfPath string, watermark string) string {
	if watermark == "" {
		watermarkVar, err := g.Config().Get(gctx.New(), "WaterMark.default")
		if err != nil {
			log.Println("获取水印配置 Error: <", err.Error(), ">")
			return ""
		}
		watermark = watermarkVar.String()
	}
	fileName := watermark + "_" + strings.Split(path.Base(pdfPath), ".")[0] + ".pdf"
	cmdStr := "/usr/local/pdfcpu watermark add -mode text -- " + "\"" + watermark + "\"" + "  \"sc:1, rot:45, mo:2,op:.3, color:.8 .8 .4\" " + pdfPath + " cache/pdf/" + fileName
	if _, ok := Doexec(cmdStr); ok {
		resultPath := "cache/pdf/" + fileName
		if PathExists(resultPath) {
			log.Printf("Convert <%s> to pdf\n", path.Base(resultPath))
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}