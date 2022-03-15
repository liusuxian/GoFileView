package cmd

import (
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/os/gcron"
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
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().Ctx,
					service.Middleware().CORS,
				)
				group.Bind(
					controller.Hello,
					controller.View,
				)
			})
			// 每天凌晨两点清理服务器文件
			_, err = gcron.Add(ctx, "0 0 2 * * *", service.ClearFile)
			if err != nil {
				logger.Error(ctx, "ClearFile Error: ", err.Error())
				return err
			}
			gres.Dump() // 打印出当前资源管理器中所有的文件列表
			s.Run()
			return nil
		},
	}
)
