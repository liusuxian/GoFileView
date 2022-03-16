package response

import (
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int    `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string `json:"message"` // 提示信息
	Data    any    `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// JsonOK 标准返回结果数据结构封装。
func JsonOK(ctx context.Context, data any) {
	Json(ctx, gcode.CodeOK.Code(), gcode.CodeOK.Message(), data)
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(ctx context.Context, code int, message string) {
	Json(ctx, code, message, nil)
	//g.Throw("exit")
}

// JsonResExit 返回JSON数据并退出当前HTTP执行函数。
func JsonResExit(ctx context.Context, code int, message string, data any) {
	Json(ctx, code, message, data)
	g.Throw("exit")
}

// Json 标准返回结果数据结构封装。
func Json(ctx context.Context, code int, message string, data any) {
	err := g.RequestFromCtx(ctx).Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if err != nil {
		logger.Error(ctx, "JsonResponse Error:", err.Error())
	}
}

// HtmlPageOK 返回 html 页面
func HtmlPageOK(ctx context.Context, data []byte) {
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(data)))
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-type:", "text/html;charset=UTF-8")
	_, err := g.RequestFromCtx(ctx).Response.Writer.Write(data)
	if err != nil {
		logger.Error(ctx, "HtmlPageOK Error:", err.Error())
	}
}

// HtmlPage 返回 html 页面
func HtmlPage(ctx context.Context, codeStr string, data []byte) {
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(len(codeStr)))
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-type:", "text/html;charset=UTF-8")
	_, err := g.RequestFromCtx(ctx).Response.Writer.Write(data)
	if err != nil {
		logger.Error(ctx, "HtmlPage Error:", err.Error())
	}
}
