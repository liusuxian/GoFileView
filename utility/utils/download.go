package utils

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// DownloadFile 下载文件
func DownloadFile(url string, localPath string) (string, error) {
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	var resp *http.Response
	var err error
	resp, err = client.Get(url)
	if err != nil {
		log.Println("DownloadFile GetHttp Error: <", err.Error(), ">")
		return "", err
	}

	var fsize int64
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		log.Println("DownloadFile GetHeader Error: <", err.Error(), "> when get file remote size")
		return "", err
	}
	if IsFileExist(localPath, fsize) {
		return "had", nil
	}

	var file *os.File
	file, err = gfile.Create(tmpFilePath)
	if err != nil {
		log.Println("DownloadFile CreateFile Error: <", err.Error(), ">")
		return "", err
	}

	defer file.Close()
	if resp.Body == nil {
		return "", gerror.New("DownloadFile Body Is Null")
	}

	var buf = make([]byte, 32*1024)
	var written int64
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, ew := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	if err == nil {
		_ = file.Close()

		var fileMd5 string
		fileMd5, err = gmd5.EncryptFile(tmpFilePath)
		if err != nil {
			log.Printf("DownloadFile GetFileMd5 <filename:%s> fail\n", gfile.Basename(localPath))
			return "", err
		}

		newPath := "cache/download/" + fileMd5 + gfile.Ext(localPath)
		err = gfile.Rename(tmpFilePath, newPath)
		if err != nil {
			log.Printf("DownloadFile FileRename <filename:%s, md5:%s> fail\n", gfile.Basename(localPath), fileMd5)
			return "", err
		}

		log.Printf("DownloadFile <filename:%s, md5:%s> success\n", gfile.Basename(localPath), fileMd5)
		return newPath, nil
	}

	return "", err
}
