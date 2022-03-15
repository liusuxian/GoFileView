package utils

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

// ConvertToPDF 转pdf
func ConvertToPDF(filePath string) string {
	// 判断转换后的pdf文件是否已经存在
	fileName := strings.Split(path.Base(filePath), ".")[0] + ".pdf"
	fileOld := "cache/pdf/" + fileName
	log.Println("ConvertToPDF filePath: ", filePath)
	log.Println("ConvertToPDF fileOld: ", fileOld)
	if FileExit(fileOld) {
		return fileOld
	}

	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		commandName = "cmd"
		params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	} else if runtime.GOOS == "linux" {
		commandName = "libreoffice"
		params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	}

	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + strings.Split(path.Base(filePath), ".")[0] + ".pdf"
		if PathExists(resultPath) {
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
	fileName := strings.Split(path.Base(filePath), ".")[0]
	fileExt := path.Ext(filePath)
	log.Println("ConvertToImg filePath: ", filePath)
	if fileExt != ".pdf" {
		return ""
	}

	// 判断转换后的jpg文件是否已经存在
	fileOld := "cache/convert/" + fileName
	if FileExit(fileOld) {
		return fileOld
	}

	if !PathExists("cache/convert/" + fileName) {
		err := os.Mkdir("cache/convert/"+fileName, os.ModePerm)
		if err != nil {
			log.Println("ConvertToImg 创建目录 Error: <", err.Error(), ">")
		}
	}

	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		commandName = "cmd"
		params = []string{"/c", "magick", "convert", "-density", "130", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	} else if runtime.GOOS == "linux" {
		commandName = "convert"
		params = []string{"-density", "150", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	}
	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/convert/" + strings.Split(path.Base(filePath), ".")[0]
		if PathExists(resultPath) {
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
	fileName := strings.Split(path.Base(filePath), ".")[0] + ".pdf"
	fileOld := "cache/pdf/" + fileName
	log.Println("ConvertToPDF filePath: ", filePath)
	log.Println("ConvertToPDF fileOld: ", fileOld)
	if FileExit(fileOld) {
		return fileOld
	}
	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		return ""
	} else if runtime.GOOS == "linux" {
		commandName = "java"
		params = []string{"-jar", "/usr/local/emailconverter-2.5.3-all.jar", filePath, "-o ", "cache/pdf/" + fileName}
	}
	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + strings.Split(path.Base(filePath), ".")[0] + ".pdf"
		if PathExists(resultPath) {
			return resultPath
		} else {
			log.Println("MsgToPdf resultPath NotExists: ", resultPath)
			return ""
		}
	} else {
		return ""
	}
}
