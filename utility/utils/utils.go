package utils

import (
	"bytes"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"os"
	"os/exec"
)

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

func IsFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	_ = gfile.Remove(filename)
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
