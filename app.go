package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/*
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: %v\n", r)
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
*/

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()  // Disable Console Color
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

	// TODO set trusted proxies to list of CIDR ranges
	r.SetTrustedProxies(nil)

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
	// Listen and Server in 0.0.0.0:8080
	log.Printf("Listening on port %v\n", port)
	r.Run(":" + port)
}
