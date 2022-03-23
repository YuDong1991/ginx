package gin

import "context"

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}
