package handlers

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	cache "github.com/Julia-Marcal/reusable-api/internal/cache"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func RateLimiting() gin.HandlerFunc {
	_, store := cache.RedisInit()

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	return mw
}
