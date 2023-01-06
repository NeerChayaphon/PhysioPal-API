package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func CacheMiddleware(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.Path
		val, err := client.Get(key).Result()
		if err == nil {
			c.String(200, val)
			c.Abort()
			return
		}

		c.Next()

		// after request
		client.Set(key, c.Writer.Status(), 0)
	}
}
