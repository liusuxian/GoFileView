package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
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

// Doexec 直接通过字符串执行shell命令，不拼接命令
func Doexec(cmdStr string) (string, bool) {
	cmd := exec.Command("bash", "-c", cmdStr)
	log.Println("Doexec cmd: ", cmd)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Doexec Error: <", err.Error(), "> when exec command read out buffer")
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

func IsFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	os.Remove(filename)
	return false
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
		log.Println("GetFileMD5 Error: <", err, "> when open file to get md5")
		return ""
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("GetFileMD5 Error: <", err, "> when get md5")
		return ""
	}
	f.Close()
	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

// GetFilenameOnly 获取路径中不带后缀的文件名
func GetFilenameOnly(filePath string) string {
	filenameWithSuffix := path.Base(filePath)
	fileSuffix := path.Ext(filenameWithSuffix)
	return strings.TrimSuffix(filenameWithSuffix, fileSuffix)
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

// 执行shell命令
func interactiveToexec(commandName string, params []string) (string, bool) {
	cmd := exec.Command(commandName, params...)
	buf, err := cmd.Output()
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if err != nil {
		log.Println("interactiveToexec Error: <", err.Error(), "> when exec command read out buffer")
		return "", false
	} else {
		return string(buf), true
	}
}
