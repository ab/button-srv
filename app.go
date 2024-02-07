package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()  // Disable Console Color
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	// Listen and serve on 0.0.0.0:8080
	log.Printf("Listening on port %v\n", port)
	r.Run(":" + port)
}
