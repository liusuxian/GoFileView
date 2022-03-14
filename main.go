package main

import (
	_ "GoFileView/internal/packed"

	"GoFileView/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
