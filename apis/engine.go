package apis

import (
	"net/http"
)

// type HandleFunc func(http.ResponseWriter, *http.Request)  v1.0
// v2.0
type HandleFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// 实现ServeHTTP()来实现net/http包中自定义handler接口，通过该接口我们能自定义处理逻辑
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 每个请求都会创建对应的上下文，并匹配对应路由进行处理
	ctx := newContext(w, r)
	e.router.handle(ctx)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 添加路由函数(调用router结构体方法)
func (e *Engine) addRoute(method string, pattern string, fn HandleFunc) {
	e.router.addRoute(method, pattern, fn)
}

func (e *Engine) GET(pattern string, fn HandleFunc) {
	e.addRoute("GET", pattern, fn)
}

func (e *Engine) POST(pattern string, fn HandleFunc) {
	e.addRoute("POST", pattern, fn)
}
