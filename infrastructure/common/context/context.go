package context

import (
	"gateway/infrastructure/common/errcode"
	"gateway/infrastructure/util/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
	Logger logging.LoggerI
}

func (c *Context) Success(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func (c *Context) Fail(code int32, msg string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func (c *Context) FailWithErrCode(err *errcode.ErrCode, data interface{}) {
	c.Fail(err.Code, err.Msg, data)
}
