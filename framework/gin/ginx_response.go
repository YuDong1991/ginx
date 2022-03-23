package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type GinxResponse interface {
	// Json输出
	XJson(obj interface{}) GinxResponse

	// Jsonp输出
	XJsonp(obj interface{}) GinxResponse

	//xml输出
	XXml(obj interface{}) GinxResponse

	// html输出
	XHtml(template string, obj interface{}) GinxResponse

	// string
	XText(format string, values ...interface{}) GinxResponse

	// 重定向
	XRedirect(path string) GinxResponse

	// header
	XSetHeader(key string, val string) GinxResponse

	// Cookie
	XSetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) GinxResponse

	// 设置状态码
	XSetStatus(code int) GinxResponse

	// 设置200状态
	XSetOkStatus() GinxResponse
}

func (c *Context) XJson(obj interface{}) GinxResponse {
	bytes, err := json.Marshal(obj)
	if err != nil {
		c.XSetStatus(http.StatusInternalServerError)
	}

	c.XSetHeader("Content-Type", "application/json")
	_, _ = c.Writer.Write(bytes)
	return c
}

func (c *Context) XJsonp(obj interface{}) GinxResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) XXml(obj interface{}) GinxResponse {
	bytes, err := xml.Marshal(obj)
	if err != nil {
		c.XSetStatus(http.StatusInternalServerError)
	}

	c.XSetHeader("Content-Type", "application/xml")
	_, _ = c.Writer.Write(bytes)
	return c
}

func (c *Context) XHtml(file string, obj interface{}) GinxResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) XText(format string, values ...interface{}) GinxResponse {
	out := fmt.Sprintf(format, values...)
	c.XSetHeader("Content-Type", "application/text")
	_, _ = c.Writer.Write([]byte(out))
	return c
}

func (c *Context) XRedirect(path string) GinxResponse {
	http.Redirect(c.Writer, c.Request, path, http.StatusMovedPermanently)
	return c
}

func (c *Context) XSetHeader(key string, val string) GinxResponse {
	c.Writer.Header().Add(key, val)
	return c
}

func (c *Context) XSetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) GinxResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) XSetStatus(code int) GinxResponse {
	c.Writer.WriteHeader(code)
	return c
}

func (c *Context) XSetOkStatus() GinxResponse {
	c.Writer.WriteHeader(http.StatusOK)
	return c
}
