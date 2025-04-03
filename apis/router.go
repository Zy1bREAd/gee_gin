package apis

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	roots   map[string]*PrefixTrieNode // 请求方式 <=> 前缀树根节点
	handler map[string]HandleFunc      // 路由表，key对应相关的handleFunc函数()
}

func newRouter() *Router {
	return &Router{
		roots:   make(map[string]*PrefixTrieNode),
		handler: make(map[string]HandleFunc),
	}
}

// 解析请求路由
func (r *Router) parsePattern(pattern string) []string {
	patternList := strings.Split(pattern, "/")

	parts := make([]string, 0)
	// 针对动态路由的:,*参数进行处理
	for _, part := range patternList {
		if part != "" {
			parts = append(parts, part)
			// 如果遇到通配符*，则代表后续的子路由都属于当前part
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由的底层实现
func (r *Router) addRoute(method string, pattern string, fn HandleFunc) {
	// v1.0通过用户传入的方法、pattern以及函数处理来在map中构建成路由表。 r.handler[routeKey] = fn
	// v2.0 判断路由表中Method是否存在路由树，否则要在前缀树插入路由节点

	parts := r.parsePattern(pattern)
	routeKey := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		// 路由不存在，针对该pattern创建一个前缀树根节点
		r.roots[method] = &PrefixTrieNode{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	fmt.Println(r.roots[method])
	// 他也没映射起来啊？？？？？
	r.handler[routeKey] = fn
	fmt.Println(r.handler[routeKey])
	fmt.Printf("[%s]Route %s add\n", method, pattern)
}

// 根据请求的 ​HTTP 方法 和 ​路径 查找匹配的路由节点，并提取路径参数。
func (r *Router) getRoute(method, path string) (*PrefixTrieNode, map[string]string) {
	fmt.Println("get Route!!!")
	searchParts := r.parsePattern(path)
	// 初始化参数容器
	params := make(map[string]string)
	// 获取对应方法的路由根节点
	root, ok := r.roots[method]
	fmt.Println("root: ", root)
	if !ok {
		// 没有找到该路由
		return nil, nil
	}
	// 从这个根节点去找，是否存在指定路由
	n := root.Search(searchParts, 0)
	fmt.Println("current node:", n)
	// 如果找到对应路由节点，则提出动态参数
	if n != nil {
		parts := r.parsePattern(n.pattern) // 解析注册时的路径模式（如 [user, :id, profile]）
		// 遍历parts,找出动态参数并提取
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] // 将动态路由参数与实际part进行映射
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/") // 将*后面请求路径剩余部分一并赋给动态路由参数
				break

			}
		}
		return n, params
	}
	return nil, nil
}

// 匹配路由map执行对应逻辑
func (r *Router) handle(ctx *Context) {
	// v1.0 组装routeKey，对handler Map中匹配后执行对应的HandleFunc
	// v2.0 获取路由以及动态参数，
	n, params := r.getRoute(ctx.Method, ctx.Path)
	fmt.Println("<handle func>:", n, params)
	// 如果找不到则会nil
	if n != nil {

		ctx.Params = params // 存储动态参数解析的值
		routeKey := ctx.Method + "-" + ctx.Path
		r.handler[routeKey](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 Not Found: %s\n", ctx.Path)
	}
}
