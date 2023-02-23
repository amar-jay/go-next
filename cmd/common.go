package main

import (
	"fmt"
	"log"

	"github.com/amar-jay/go-api-boilerplate/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env files. $PORT must be set")
	}
}

func main() {
	fmt.Println("Starting server...")
	err := router.SetTrustedProxies([]string{"192.168.1.2", "::1"})

	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.GinContextToMiddleWare())

}
