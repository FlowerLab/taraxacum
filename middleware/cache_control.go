package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Etag", `W/"5c329-dZi4o7Cn7JCq7pOA7Ih5uhDGOGI"`)
		c.Header("Cache-Control", "max-age=2592000")
		if match := c.GetHeader("If-None-Match"); match != "" {
			c.AbortWithStatus(http.StatusNotModified)
		}
	}
}
