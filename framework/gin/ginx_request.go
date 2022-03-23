package gin

import (
	"github.com/spf13/cast"
	"mime/multipart"
)

// 代表请求包含的方法
type Request interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	XParam(key string) interface{}

	// DefaultForm表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (c *Context) QueryAll() map[string][]string {
	c.initQueryCache()
	return c.queryCache
}

func (c *Context) DefaultQueryInt(key string, def int) (int, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToInt(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToInt64(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToBool(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryString(key string, def string) (string, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return vals[0], true
		}
	}

	return def, false
}

func (c *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	DefaultParams := c.QueryAll()
	if vals, ok := DefaultParams[key]; ok {
		return vals, true
	}

	return def, false
}

func (c *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToInt(val), true
	}

	return def, false
}

func (c *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToInt64(val), true
	}

	return def, false
}

func (c *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToFloat32(val), true
	}

	return def, false
}

func (c *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToFloat64(val), true
	}

	return def, false
}

func (c *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToBool(val), true
	}

	return def, false
}

func (c *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := c.XParam(key); val != nil {
		return cast.ToString(val), true
	}

	return def, false
}

func (c *Context) XParam(key string) interface{} {
	if val, ok := c.params.Get(key); ok {
		return val
	}

	return nil
}

func (c *Context) FormAll() map[string][]string {
	c.initFormCache()
	return c.formCache
}

func (c *Context) DefaultFormInt(key string, def int) (int, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToInt(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToInt64(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return cast.ToBool(vals[0]), true
		}
	}

	return def, false
}

func (c *Context) DefaultFormString(key string, def string) (string, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return vals[0], true
		}
	}

	return def, false
}

func (c *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		return vals, true
	}

	return def, false
}

func (c *Context) FromFile(key string) (*multipart.FileHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Context) DefaultForm(key string) interface{} {
	DefaultParams := c.FormAll()
	if vals, ok := DefaultParams[key]; ok {
		if len(vals) >= 0 {
			return vals[0]
		}
	}

	return nil
}
