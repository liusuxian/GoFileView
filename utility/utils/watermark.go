package utils

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

// WaterMark pdf增加水印
func WaterMark(ctx context.Context, pdfPath string, watermark string) (string, error) {
	if watermark == "" {
		watermarkVar, err := g.Cfg().Get(ctx, "WaterMark.default", "liusuxian")
		if err != nil {
			return "", err
		}
		watermark = watermarkVar.String()
	}

	fileName := watermark + "_" + gfile.Name(pdfPath) + ".pdf"
	//cmdStr := "/usr/local/pdfcpu watermark add -mode text -- " + "\"" + watermark + "\"" + "  \"rot:45, mo:2, op:.3, color:.8 .8 .4\" " + "\"" + pdfPath + "\" " + "\"cache/pdf/" + fileName + "\""
	cmdStr := "/usr/local/pdfcpu stamp add -mode text -- " + "\"" + watermark + "\"" + "  \"rot:45, mo:2, op:.3, color:.8 .8 .4\" " + "\"" + pdfPath + "\" " + "\"cache/pdf/" + fileName + "\""
	err := Doexec(cmdStr)
	if err != nil {
		return "", err
	}
	resultPath := "cache/pdf/" + fileName
	if gfile.Exists(resultPath) {
		return resultPath, nil
	} else {
		return "", gerror.Newf("WaterMark resultPath NotExists: %s", resultPath)
	}
}
