package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.DisableConsoleColor()  // Disable Console Color
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

	// Never use X-Forwarded-For
	r.SetTrustedProxies(nil)

	// Get client IP from fly.io header
	// https://fly.io/docs/reference/runtime-environment/#fly-client-ip
	r.TrustedPlatform = "Fly-Client-IP"

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "region": os.Getenv("FLY_REGION")})
	})

	// Root default page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"region":   os.Getenv("FLY_REGION"),
			"clientIP": c.ClientIP(),
		})
	})

	return r
}
