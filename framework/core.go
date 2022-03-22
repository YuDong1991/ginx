package framework

import (
	"net/http"
)

// Core represent core struct
type Core struct {
	Router
}

func NewCore() *Core {
	return &Core{
		Router: NewDynamicRouter(),
	}
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 封装自定义 Context
	ctx := NewContext(request, writer)

	// 寻找路由
	handler := c.FindRouteByRequest(request)
	if handler == nil {
		_ = ctx.Json(http.StatusNotFound, "not found")
		return
	}

	// 调用处理器
	if err := handler(ctx); err != nil {
		_ = ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}

func (c *Core) RouteGroup(prefix string) RouteGroup {
	return newSimpleRouteGroup(prefix, c)
}
