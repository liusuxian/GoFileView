package utils

import (
	"bytes"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"os"
	"os/exec"
)

// Doexec 直接通过字符串执行shell命令，不拼接命令
func Doexec(cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	log.Println("Doexec cmd: ", cmd)
	_, err := cmd.CombinedOutput()
	return err
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
func interactiveToexec(commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	log.Println("interactiveToexec cmd: ", cmd)
	_, err := cmd.Output()
	cmd.Stderr = bytes.NewBuffer(nil)
	return err
}
