package service

import (
	"GoFileView/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

var insMiddleware = sMiddleware{}

// Middleware 中间件管理服务
func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(req *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: req.Session,
	}
	Context().Init(req, customCtx)
	// 执行下一步请求逻辑
	req.Middleware.Next()
}

// CORS 允许接口跨域请求
func (s *sMiddleware) CORS(req *ghttp.Request) {
	req.Response.CORSDefault()
	req.Middleware.Next()
}
