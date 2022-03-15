package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gres"

	"GoFileView/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.View,
				)
			})
			gres.Dump() // 打印出当前资源管理器中所有的文件列表
			s.Run()
			return nil
		},
	}
)
