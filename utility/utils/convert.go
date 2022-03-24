package utils

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"runtime"
)

// ConvertToPDF 转pdf
func ConvertToPDF(filePath string) (string, error) {
	// 判断转换后的pdf文件是否已经存在
	fileName := gfile.Name(filePath) + ".pdf"
	fileOld := "cache/pdf/" + fileName
	if gfile.Exists(fileOld) {
		return fileOld, nil
	}

	commandName := ""
	var params []string
	switch runtime.GOOS {
	case "windows":
		commandName = "cmd"
		params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	case "linux":
		commandName = "libreoffice"
		params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	case "darwin":
		commandName = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
		params = []string{"--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	default:
		return "", gerror.Newf("ConvertToPDF Nonsupport OS: %s", runtime.GOOS)
	}

	err := interactiveToexec(commandName, params)
	if err != nil {
		return "", err
	}
	resultPath := "cache/pdf/" + gfile.Name(filePath) + ".pdf"
	if gfile.Exists(resultPath) {
		return resultPath, nil
	} else {
		return "", gerror.Newf("ConvertToPDF resultPath NotExists: : %s", resultPath)
	}
}

// ConvertToImg 转图片
func ConvertToImg(filePath string) (string, error) {
	fileName := gfile.Name(filePath)
	fileExt := gfile.Ext(filePath)
	if fileExt != ".pdf" {
		return "", gerror.Newf("ConvertToImg Nonsupport FileType: %s", fileExt)
	}

	// 判断转换后的jpg文件是否已经存在
	fileOld := "cache/convert/" + fileName
	if gfile.Exists(fileOld) {
		return fileOld, nil
	}

	if !gfile.Exists("cache/convert/" + fileName) {
		err := gfile.Mkdir("cache/convert/" + fileName)
		if err != nil {
			return "", err
		}
	}

	commandName := ""
	var params []string
	switch runtime.GOOS {
	case "windows":
		commandName = "cmd"
		params = []string{"/c", "magick", "convert", "-density", "130", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	case "linux":
		commandName = "convert"
		params = []string{"-density", "150", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	case "darwin":
		commandName = "convert"
		params = []string{"-density", "150", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	default:
		return "", gerror.Newf("ConvertToImg Nonsupport OS: %s", runtime.GOOS)
	}

	err := interactiveToexec(commandName, params)
	if err != nil {
		return "", err
	}
	resultPath := "cache/convert/" + gfile.Name(filePath)
	if gfile.Exists(resultPath) {
		return resultPath, nil
	} else {
		return "", gerror.Newf("ConvertToImg resultPath NotExists: : %s", resultPath)
	}
}

// MsgToPdf 只支持linux
func MsgToPdf(filePath string) (string, error) {
	// 判断转换后的pdf文件是否已经存在
	fileName := gfile.Name(filePath) + ".pdf"
	fileOld := "cache/pdf/" + fileName
	if gfile.Exists(fileOld) {
		return fileOld, nil
	}

	commandName := ""
	var params []string
	switch runtime.GOOS {
	case "linux":
		commandName = "java"
		params = []string{"-jar", "/usr/local/emailconverter-2.5.3-all.jar", filePath, "-o ", "cache/pdf/" + fileName}
	default:
		return "", gerror.Newf("MsgToPdf Nonsupport OS: %s", runtime.GOOS)
	}

	err := interactiveToexec(commandName, params)
	if err != nil {
		return "", err
	}
	resultPath := "cache/pdf/" + gfile.Name(filePath) + ".pdf"
	if gfile.Exists(resultPath) {
		return resultPath, nil
	} else {
		return "", gerror.Newf("MsgToPdf resultPath NotExists: : %s", resultPath)
	}
}
