package middleware

import (
	"gateway/infrastructure/common/errcode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimit(rate int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(time.Duration(1e9/rate), rate)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) <= 0 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": errcode.LibRateLimit,
				"msg":  errcode.LibRateLimit.String(),
				"data": nil,
			})
			return
		}
		c.Next()
	}
}
