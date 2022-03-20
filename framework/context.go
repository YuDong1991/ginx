package framework

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	// 是否超时标记位
	hasTimeout bool
	// 写保护锁
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		handler:        nil,
		hasTimeout:     false,
		writerMux:      &sync.Mutex{},
	}
}

// #基本功能方法

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHandler(handler ControllerHandler) {
	ctx.handler = handler
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// #实现标准库 Context 接口

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// #获取 url 查询数据

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len := len(vals); len > 0 {
			if intVal, err := strconv.Atoi(vals[len-1]); err == nil {
				return intVal
			}
		}
	}
	return def
}

func (ctx Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len := len(vals); len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// #获取 post 表单数据

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len := len(vals); len > 0 {
			if intVal, err := strconv.Atoi(vals[len-1]); err == nil {
				return intVal
			}
		}
	}
	return def
}

func (ctx Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len := len(vals); len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

// #获取 post application/json

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request == nil {
		return errors.New("ctx.request empty")
	}

	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, obj)
}

// # response 方法

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.hasTimeout {
		return nil
	}

	ctx.responseWriter.Header().Set("Content", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = ctx.responseWriter.Write(byt)
	return err
}

func (ctx *Context) Html(status int, obj interface{}) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
