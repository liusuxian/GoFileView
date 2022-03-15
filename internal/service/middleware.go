package service

import (
	"GoFileView/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

var (
	insMiddleware = sMiddleware{
		LoginUrl: "/login",
	}
)

// Middleware 中间件管理服务
func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context().Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// CORS 允许接口跨域请求
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
