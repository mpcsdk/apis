package cmd

import (
	"apis/internal/controller/chaindata"
	"apis/internal/controller/enhanced"
	"apis/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"go.opentelemetry.io/otel/trace"
)

var apiLimit = mpccode.CodeApiLimit

func RateLimit(r *ghttp.Request) {
	if service.RateLimiter().Allow() {
		r.Middleware.Next()
	} else {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    100,
			Message: "api limit",
			Data:    nil,
		})
	}

}
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.ClearBuffer()

		///
		spanCtx := trace.SpanContextFromContext(r.Context())
		traceId := spanCtx.TraceID()
		///
		code := gcode.CodeInternalError
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: code.Message(),
			Data:    traceId.String(),
		})
	}
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func ResponseHandler(r *ghttp.Request) {
	g.Log().Info(r.Context(), "Request:", r.GetUrl(), r.GetBodyString())
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	r.SetError(nil)
	if code == gcode.CodeNil {
		if err != nil {
			code = gcode.CodeInternalError
		} else {
			code = gcode.CodeOK
		}
	}
	g.Log().Info(r.Context(), "Response:", r.GetUrl(), res)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data: func() interface{} {
			detail := code.Detail()
			if detail != nil {
				return detail
			} else {
				return res
			}
		}(),
	})
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(RateLimit)
				group.Middleware(MiddlewareErrorHandler)
				group.Middleware(MiddlewareCORS)
				group.Middleware(ResponseHandler)
				group.Group("/chaindata", func(group *ghttp.RouterGroup) {
					group.Bind(chaindata.NewV1())
				})
				group.Group("/enhanced", func(group *ghttp.RouterGroup) {
					group.Bind(enhanced.NewV1())
				})
			})
			s.Run()
			return nil
		},
	}
)
