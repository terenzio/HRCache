package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CacheMiddleware(cache *Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			key := c.Request.RequestURI
			if data, found := cache.Get(key); found {
				c.JSON(http.StatusOK, data)
				c.Abort()
				return
			}
			c.Next()
			cache.Set(key, c.Writer.Status())
		} else {
			c.Next()
			cache.Delete(c.Param("id"))
		}
	}
}
