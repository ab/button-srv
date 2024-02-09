package main

import (
	"log"
	"os"

	"github.com/ab/button-srv/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := server.InitRouter()
	// Listen and serve on 0.0.0.0:8080
	log.Printf("button-srv/gin listening on port %v\n", port)
	r.Run(":" + port)
}
