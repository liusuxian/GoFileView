package service

import (
	"GoFileView/utility/logger"
	"GoFileView/utility/utils"
	"context"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"time"
)

type NowFile struct {
	Md5            string
	Ext            string
	LastActiveTime int64
}

var (
	AllFile      map[string]*NowFile
	AllOfficeEtx = []string{".doc", ".docx", ".xls", ".xlsx", ".csv", ".ppt", ".pptx", ".txt", ".msg"}
	AllImageEtx  = []string{".jpg", ".png", ".gif"}
)

func OfficePage(imgPath string) []byte {
	rd, _ := ioutil.ReadDir(imgPath)
	dataByte := gfile.GetBytes("resource/public/html/office.html")
	dataStr := string(dataByte)
	htmlCode := ""
	for i := 0; i < len(rd); i++ {
		htmlCode = htmlCode +
			`<img class="my-photo" alt="loading" title="查看大图" style="cursor: pointer;"
		 data-src="/view/office?url=` + gfile.Basename(imgPath) + "/" + gconv.String(i) + ".jpg" +
			`" src="/static/image/loading.gif" ">`

	}
	dataStr = gstr.Replace(dataStr, "{{AllImages}}", htmlCode, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func ImagePage(filePath string) []byte {
	dataByte := gfile.GetBytes("resource/public/html/image.html")
	dataStr := string(dataByte)
	imageUrl := "/view/img?url=" + gfile.Basename(filePath)
	htmlCode := `<li>
					<img id="` + imageUrl + `" url="` + imageUrl + `"
						src="` + imageUrl + `" width="1px" height="1px">
				 </li>`
	dataStr = gstr.Replace(dataStr, "{{AllImages}}", htmlCode, -1)
	dataStr = gstr.Replace(dataStr, "{{FirstPath}}", imageUrl, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func PdfPage(filePath string) []byte {
	dataByte := gfile.GetBytes("resource/public/html/pdf.html")
	dataStr := string(dataByte)

	pdfUrl := "/view/pdf?url=" + gfile.Basename(filePath)
	dataStr = gstr.Replace(dataStr, "{{url}}", pdfUrl, -1)

	dataByte = []byte(dataStr)
	return dataByte
}

func PdfPageDownload(filePath string) []byte {
	dataByte := gfile.GetBytes("resource/public/html/pdf.html")
	dataStr := string(dataByte)
	pdfUrl := "/view/img?url=" + gfile.Basename(filePath)
	dataStr = gstr.Replace(dataStr, "{{url}}", pdfUrl, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func MdPage(filepath string) []byte {
	fileByte := gfile.GetBytes(filepath)
	unsafe := blackfriday.MarkdownCommon(fileByte)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	dataByte := gfile.GetBytes("resource/public/html/md.html")
	dataStr := string(dataByte)

	dataStr = gstr.Replace(dataStr, "{{url}}", string(html), -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func IsHave(fileName string) bool {
	fileName = gfile.Name(fileName)
	if _, ok := AllFile[fileName]; ok {
		AllFile[fileName].LastActiveTime = time.Now().Unix()
		return true
	} else {
		return false
	}
}

func SetFileMap(fileName string) {
	ext := gfile.Ext(fileName)
	fileName = gfile.Name(fileName)
	if _, ok := AllFile[fileName]; ok {
		AllFile[fileName].LastActiveTime = time.Now().Unix()
		return
	} else {
		temp := &NowFile{
			Md5:            fileName,
			Ext:            ext,
			LastActiveTime: time.Now().Unix(),
		}
		AllFile[fileName] = temp
	}
}

// ClearFile 清除目录文件
func ClearFile(ctx context.Context) {
	logger.Info(ctx, "-------------开始清除服务器文件------------")
	// 删除图片目录里的所有文件
	dir1, err := ioutil.ReadDir("cache/convert")
	if err != nil {
		logger.Error(ctx, "ClearFile 读取<cache/convert>目录错误: ", err.Error())
		return
	}
	for _, d := range dir1 {
		_ = gfile.Remove(gfile.Join([]string{"cache/convert", d.Name()}...))
	}
	logger.Info(ctx, "cache/convert 已清除")

	// 删除本地下载文件目录里的所有文件
	dir2, err := ioutil.ReadDir("cache/download")
	if err != nil {
		logger.Error(ctx, "ClearFile 读取<cache/download>目录错误: ", err.Error())
		return
	}
	for _, d := range dir2 {
		_ = gfile.Remove(gfile.Join([]string{"cache/download", d.Name()}...))
	}
	logger.Info(ctx, "cache/download 已清除")

	// 删除pdf目录里的所有文件
	dir3, err := ioutil.ReadDir("cache/pdf")
	if err != nil {
		logger.Error(ctx, "ClearFile 读取<cache/pdf>目录错误: ", err.Error())
		return
	}
	for _, d := range dir3 {
		_ = gfile.Remove(gfile.Join([]string{"cache/pdf", d.Name()}...))
	}
	logger.Info(ctx, "cache/pdf 已清除")
	logger.Info(ctx, "---------------清除文件已完成--------------")
}

func GetAllFile(pathname string) ([]map[string]string, error) {
	var s []map[string]string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return s, err
	}

	for _, fi := range rd {
		tmp := map[string]string{}
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			tmp["path"] = fullName
			tmp["name"] = fi.Name()
			tmp["type"] = gfile.Ext(fullName)
		}
		s = append(s, tmp)
	}
	return s, nil
}

// ExcelPage 将Excel转html
func ExcelPage(filePath string) []byte {
	ret := utils.ExcelParse(filePath)
	html := `
			<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN"><html><head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8"/>
			<title></title>	<style type="text/css">body,div,table,thead,tbody,tfoot,tr,th,td,p { font-family:"等线"; font-size:x-small }
			a.comment-indicator:hover + comment { background:#ffd; position:absolute; display:block; border:1px solid black; padding:0.5em;  }
			a.comment-indicator { background:red; display:inline-block; border:1px solid black; width:0.5em; height:0.5em;  }
			comment { display:none;  } 	</style>
	`
	html += "<p><center>		<h1>Overview</h1>"
	for i := 0; i < len(ret); i++ {
		html += "<A HREF=\"#table" + gconv.String(i) + "\" style = \"font-size: 30px;\"  >Sheet" + gconv.String(i+1) + " </A><br>"
	}
	html += "</center></p><hr>"

	for k, v := range ret {
		html += "<A NAME=\"table" + gconv.String(k) + "\"   style = \"color: #337ab7;\">"
		html += "<h1>Sheet" + gconv.String(k+1) + "</h1></A>"
		html += `
			<table  class = "table table-striped" cellspacing ="0" border ="0"  style= "width: 100%;max-width: 100%;"> 
		`
		for _, vs := range gconv.SliceAny(v["title"]) {
			num := len(gconv.String(vs)) * 10
			html += "<colgroup width=\"" + gconv.String(num) + "\"></colgroup>  "
		}

		html += "<tr>"
		for _, vs := range gconv.SliceAny(v["title"]) {
			html += "<td height=\"19\" align=\"left\" valign=bottom><font color=\"#000000\">" + gconv.String(vs) + "</font></td>	 "
		}
		html += "</tr>"

		for _, vs := range gconv.SliceAny(v["resourceArr"]) {
			html += "<tr>"
			for _, vss := range gconv.SliceAny(vs) {
				html += "<td height=\"19\" align=\"left\" valign=bottom><font color=\"#000000\">" + gconv.String(vss) + "</font></td>	 "
			}
			html += "</tr>"
		}
		html += "</table>"

	}
	html += `
		</html>
		<script src="/html/js/jquery-3.0.0.min.js" type="text/javascript">
		</script><script src="/html/js/excel.header.js" type="text/javascript">
		</script><link rel="stylesheet" href="/html/css/bootstrap.min.css">
		`
	dataByte := gfile.GetBytes("resource/public/html/excel.html")
	dataStr := string(dataByte)

	dataStr = gstr.Replace(dataStr, "{{url}}", html, -1)
	dataByte = []byte(dataStr)
	return dataByte
}
