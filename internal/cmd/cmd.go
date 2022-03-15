package cmd

import (
	"GoFileView/internal/controller"
	"GoFileView/internal/service"
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			serverRoot := g.Cfg().MustGet(ctx, "server.serverRoot", "resource/public").String()
			s.AddStaticPath("/static", serverRoot+"/resource")
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					controller.Hello,
				)
				group.Group("/view", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.View,
						controller.Img,
						controller.Pdf,
						controller.Office,
						controller.Upload,
						controller.Delete,
					)
				})
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
