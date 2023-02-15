package server

import (
	"github.com/gin-gonic/gin"
	ginHelper "github.com/livegoplayer/go_gin_helper"
	"github.com/livegoplayer/go_helper/utils"
)

// MustParams 用于定义存放请求中的公共参数
type MustParams struct {
}

func (p MustParams) ToMap() map[string]interface{} {
	return utils.ToMap(p)
}

type Controller func(ctx *Context)

type FormErr struct {
	reason string
}

/*
继承自gin.Context的类，用于简化调用, 优化结构。
类中的数据不会在controller之间传递，要保存的数据，还是要通过gin的方法保存
调用NewCtl可以根据新的controller生成gin使用过的
*/
type Context struct {
	*gin.Context
	Param *MustParams // 中间件中解析的公共参数结构体
}

func (c *Context) SuccessResp(objs ...interface{}) {
	ginHelper.SuccessResp("ok", objs...)
}

func (c *Context) ErrorResp(code int, msg string, objs ...interface{}) {
	ginHelper.ErrorResp(code, msg, objs...)
}

func (c *Context) CheckError(error error, msg ...string) {
	ginHelper.CheckError(error, msg...)
}

func (c *Context) AuthResp(msg string, url string) {
	ginHelper.AuthResp(msg, url)
}

// NewCtl 根据自定义的Context的controller, 生成gin使用的controller
func NewCtl(controller Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		pm := &MustParams{}

		p, isExist := c.Get("params")
		if isExist {
			//p = &MustParams{}
			m := *(p.(*interface{}))
			pm = m.(*MustParams)
		}
		ctx := &Context{c, pm}

		controller(ctx)
	}
}
