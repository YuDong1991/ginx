package framework

import (
	"log"
	"net/http"
	"strings"
)

type RouteRegister interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type RouteMatching interface {
	FindRouteByRequest(request *http.Request) ControllerHandler
}

type Router interface {
	RouteRegister
	RouteMatching
}

// #不带动态路由功能的简单路由器

type SimpleRouter map[string]map[string]ControllerHandler

func NewSimpleRouter() Router {
	return SimpleRouter{
		http.MethodGet:    {},
		http.MethodPost:   {},
		http.MethodDelete: {},
		http.MethodPut:    {},
	}
}

// #实现路由器接口

func (c SimpleRouter) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c[http.MethodGet][upperUrl] = handler
}

func (c SimpleRouter) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c[http.MethodPost][upperUrl] = handler
}

func (c SimpleRouter) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c[http.MethodPut][upperUrl] = handler
}

func (c SimpleRouter) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c[http.MethodDelete][upperUrl] = handler
}

func (c SimpleRouter) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c[upperMethod]; ok {
		return methodHandlers[upperUri]
	}
	return nil
}

// #带动态路由功能的路由器

type DynamicRouter map[string]*Tree

func NewDynamicRouter() Router {
	return &DynamicRouter{
		http.MethodGet:    &Tree{root: new(node)},
		http.MethodPost:   &Tree{root: new(node)},
		http.MethodDelete: &Tree{root: new(node)},
		http.MethodPut:    &Tree{root: new(node)},
	}
}

func (d DynamicRouter) Get(s string, handler ControllerHandler) {
	if err := d[http.MethodGet].AddRoute(s, handler); err != nil {
		log.Fatal("register route err: ", err)
	}
}

func (d DynamicRouter) Post(s string, handler ControllerHandler) {
	if err := d[http.MethodPost].AddRoute(s, handler); err != nil {
		log.Fatal("register route err: ", err)
	}
}

func (d DynamicRouter) Put(s string, handler ControllerHandler) {
	if err := d[http.MethodPut].AddRoute(s, handler); err != nil {
		log.Fatal("register route err: ", err)
	}
}

func (d DynamicRouter) Delete(s string, handler ControllerHandler) {
	if err := d[http.MethodDelete].AddRoute(s, handler); err != nil {
		log.Fatal("register route err: ", err)
	}
}

func (d DynamicRouter) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	if tree, ok := d[upperMethod]; ok {
		return tree.FindHandler(upperUri)
	}

	return nil
}

// #分组路由注册

type RouteGroup interface {
	RouteRegister
	RouteGroup(string) RouteGroup
}

type SimpleRouteGroup struct {
	router Router
	prefix string
}

func newSimpleRouteGroup(prefix string, router Router) RouteGroup {
	return SimpleRouteGroup{
		router: router,
		prefix: prefix,
	}
}

func (c SimpleRouteGroup) RouteGroup(s string) RouteGroup {
	c.prefix += s
	return c
}

func (c SimpleRouteGroup) Get(uri string, handler ControllerHandler) {
	uri = c.prefix + uri
	upperUri := strings.ToUpper(uri)
	c.router.Get(upperUri, handler)
}

func (c SimpleRouteGroup) Post(uri string, handler ControllerHandler) {
	uri = c.prefix + uri
	upperUri := strings.ToUpper(uri)
	c.router.Post(upperUri, handler)
}

func (c SimpleRouteGroup) Put(uri string, handler ControllerHandler) {
	uri = c.prefix + uri
	upperUri := strings.ToUpper(uri)
	c.router.Put(upperUri, handler)
}

func (c SimpleRouteGroup) Delete(uri string, handler ControllerHandler) {
	uri = c.prefix + uri
	upperUri := strings.ToUpper(uri)
	c.router.Delete(upperUri, handler)
}
