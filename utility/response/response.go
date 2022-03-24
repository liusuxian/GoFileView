package response

import (
	"GoFileView/utility/logger"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

// HtmlText 返回 HtmlText
func HtmlText(ctx context.Context, size int, data []byte) {
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-length", strconv.Itoa(size))
	g.RequestFromCtx(ctx).Response.Writer.Header().Set("content-type:", "text/html;charset=UTF-8")
	_, err := g.RequestFromCtx(ctx).Response.Writer.Write(data)
	if err != nil {
		logger.Error(ctx, "HtmlPage Error:", err.Error())
	}
}
