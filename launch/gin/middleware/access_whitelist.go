package middleware

import (
	"gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

// AccessWhitelist 访问白名单
func AccessWhitelist(c *context.Context) {
	var (
		enabled = viper.GetBool("access_whitelist.enabled")
	)

	var (
		ipList = viper.GetStringSlice("access_whitelist.ip_list")
		userIP = getRequestIP(c)
	)

	if userIP != "" && sliceContain(ipList, userIP) {
		c.Next()
		return
	}

	c.FailWithErrCode(errcode.LibNotInWhitelist, gin.H{
		"ip":      userIP,
		"enabled": enabled,
	})
	c.Abort()
}

// getRequestIP 获取ip
func getRequestIP(c *context.Context) string {
	original := c.GetHeader("X-Original-Forwarded-For")
	if original != "" {
		ips := strings.Split(original, ",")
		return ips[0]
	}
	if c.GetHeader("X-Forwarded-For") != "" {
		return c.GetHeader("X-Forwarded-For")
	}

	return c.RemoteIP()
}

func sliceContain[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
