package apis

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	handler map[string]HandleFunc // 路由表，对应相关的handleFunc函数()
}

func newRouter() *Router {
	return &Router{
		handler: map[string]HandleFunc{},
	}
}

// 添加路由的底层实现
func (r *Router) addRoute(method string, pattern string, fn HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	routeKey := fmt.Sprintf("%s-%s", method, pattern)
	r.handler[routeKey] = fn
}

// 匹配路由map执行对应逻辑
func (r *Router) handle(ctx *Context) {
	routeKey := ctx.Method + "-" + ctx.Path
	fmt.Println("test:", r.handler, routeKey)
	if fn, ok := r.handler[routeKey]; ok {
		fn(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
