package middleware

import (
	"gateway/infrastructure/util/def"
	"gateway/infrastructure/util/util/random"
	"github.com/gin-gonic/gin"
)

func SetV(c *gin.Context) {
	xRequestID := c.GetHeader(def.XRequestID)
	if xRequestID == "" {
		xRequestID = random.UUID()
	}
	c.Set(def.XRequestID, xRequestID)
	c.Set("v1", c.Request.URL.Path)
	c.Next()
}
