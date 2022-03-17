package utils

import (
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"runtime"
)

// ConvertToPDF 转pdf
func ConvertToPDF(filePath string) string {
	// 判断转换后的pdf文件是否已经存在
	fileName := gfile.Name(filePath) + ".pdf"
	fileOld := "cache/pdf/" + fileName
	log.Println("ConvertToPDF filePath: ", filePath)
	log.Println("ConvertToPDF fileOld: ", fileOld)
	if gfile.Exists(fileOld) {
		return fileOld
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
		log.Println("ConvertToPDF Nonsupport OS: ", runtime.GOOS)
		return ""
	}

	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + gfile.Name(filePath) + ".pdf"
		if gfile.Exists(resultPath) {
			return resultPath
		} else {
			log.Println("ConvertToPDF resultPath NotExists: ", resultPath)
			return ""
		}
	} else {
		return ""
	}
}

// ConvertToImg 转图片
func ConvertToImg(filePath string) string {
	fileName := gfile.Name(filePath)
	fileExt := gfile.Ext(filePath)
	log.Println("ConvertToImg filePath: ", filePath)
	if fileExt != ".pdf" {
		return ""
	}

	// 判断转换后的jpg文件是否已经存在
	fileOld := "cache/convert/" + fileName
	log.Println("ConvertToImg fileOld: ", fileOld)
	if gfile.Exists(fileOld) {
		return fileOld
	}

	if !gfile.Exists("cache/convert/" + fileName) {
		err := gfile.Mkdir("cache/convert/" + fileName)
		if err != nil {
			log.Println("ConvertToImg 创建目录 Error: <", err.Error(), ">")
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
		log.Println("ConvertToImg Nonsupport OS: ", runtime.GOOS)
		return ""
	}

	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/convert/" + gfile.Name(filePath)
		if gfile.Exists(resultPath) {
			return resultPath
		} else {
			log.Println("ConvertToImg resultPath NotExists: ", resultPath)
			return ""
		}
	} else {
		return ""
	}
}

// MsgToPdf 只支持linux
func MsgToPdf(filePath string) string {
	// 判断转换后的pdf文件是否已经存在
	fileName := gfile.Name(filePath) + ".pdf"
	fileOld := "cache/pdf/" + fileName
	log.Println("ConvertToPDF filePath: ", filePath)
	log.Println("ConvertToPDF fileOld: ", fileOld)
	if gfile.Exists(fileOld) {
		return fileOld
	}

	commandName := ""
	var params []string
	switch runtime.GOOS {
	case "linux":
		commandName = "java"
		params = []string{"-jar", "/usr/local/emailconverter-2.5.3-all.jar", filePath, "-o ", "cache/pdf/" + fileName}
	default:
		log.Println("MsgToPdf Nonsupport OS: ", runtime.GOOS)
		return ""
	}

	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + gfile.Name(filePath) + ".pdf"
		if gfile.Exists(resultPath) {
			return resultPath
		} else {
			log.Println("MsgToPdf resultPath NotExists: ", resultPath)
			return ""
		}
	} else {
		return ""
	}
}
