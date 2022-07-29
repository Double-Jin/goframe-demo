package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"go.opentelemetry.io/otel/trace"
	"goframe/api/v1/admin"
	"goframe/internal/common"
)

type sMiddleware struct{}

var (
	// insMiddleware is the instance of service Middleware.
	insMiddleware = sMiddleware{}
)

// Middleware returns the interface of Middleware service.
func Middleware() *sMiddleware {
	return &insMiddleware
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {

	spanCtx := trace.SpanContextFromContext(r.Context())

	reqId := guid.S(grand.B(64))
	if traceId := spanCtx.TraceID(); traceId.IsValid() {
		reqId = traceId.String()
	}

	customCtx := &admin.Context{
		Data:    make(g.Map),
		Request: r,
		ReqId:   reqId,
	}

	common.Context.Init(r, customCtx)

	r.Middleware.Next()
}


func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
