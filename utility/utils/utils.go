package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"
)

func ComparePath(a string, b string) bool {
	if len(a) >= len(b) {
		if strings.Compare(a[0:len(b)], b) == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// ConvertToPDF 转pdf
func ConvertToPDF(filePath string) string {
	// 判断转换后的pdf文件是否已经存在
	fileName := strings.Split(path.Base(filePath), ".")[0] + ".pdf"
	fileOld := "cache/pdf/" + fileName
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
		log.Println("filePath: ", filePath)
		params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	}

	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + strings.Split(path.Base(filePath), ".")[0] + ".pdf"
		if PathExists(resultPath) {
			log.Printf("Convert <%s> to pdf\n", path.Base(filePath))
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

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

// ConvertToImg 转图片
func ConvertToImg(filePath string) string {
	fileName := strings.Split(path.Base(filePath), ".")[0]
	fileExt := path.Ext(filePath)
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
			log.Println("创建目录 Error: <", err.Error(), ">")
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
			log.Printf("Convert <%s> to images\n", path.Base(filePath))
			log.Println("resultPath: ", resultPath)
			return resultPath
		} else {
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
			log.Printf("Convert <%s> to pdf\n", path.Base(filePath))
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// Doexec 直接通过字符串执行shell命令，不拼接命令
func Doexec(cmdStr string) (string, bool) {
	cmd := exec.Command("bash", "-c", cmdStr)
	log.Println("cmd: ", cmd)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error: <", err.Error(), "> when exec command read out buffer")
		return "", false
	} else {
		return string(buf), true
	}
}

// 执行shell命令
func interactiveToexec(commandName string, params []string) (string, bool) {
	cmd := exec.Command(commandName, params...)
	log.Println("cmd: ", cmd)
	buf, err := cmd.Output()
	log.Println(string(buf), err)
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if err != nil {
		log.Println("Error: <", err.Error(), "> when exec command read out buffer")
		return "", false
	} else {
		return string(buf), true
	}
}

func FileExit(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetFileMD5(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("Error: <", err, "> when open file to get md5")
		return ""
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Error: <", err, "> when get md5")
		return ""
	}
	f.Close()
	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	byteSlice := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		byteSlice[i] = byte(b)
	}
	return string(byteSlice)
}

func IsInArr(key string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			return true
		}
	}
	return false
}

// ExcelParse excel解析
func ExcelParse(filePath string) []map[string]interface{} {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Println("ExcelParseError: <", err.Error(), ">")
	}
	var resData []map[string]interface{}

	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		tmp := map[string]interface{}{}
		//遍历每一行
		var title []string
		var resourceArr [][]string
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			if rowIndex == 0 {
				for _, cell := range row.Cells {
					text := cell.String()
					title = append(title, text)
				}
				continue
			}
			//遍历每一个单元
			var result []string
			for _, cell := range row.Cells {
				text := cell.String()
				result = append(result, text)
			}
			resourceArr = append(resourceArr, result)
		}

		tmp["title"] = title
		tmp["resourceArr"] = resourceArr

		resData = append(resData, tmp)
	}
	return resData
}
