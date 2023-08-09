package middleware

import (
	"fmt"
	"gateway/infrastructure/util/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var errMsg string
			switch err := err.(type) {
			case error:
				errMsg = fmt.Sprintf("%+v", errors.WithStack(err))
			default:
				errMsg = fmt.Sprintf("%v", err)
			}
			logging.FromContext(c).Fatal("服务器内部错误", "", errMsg)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
