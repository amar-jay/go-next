package cmd 

import (
	"log"

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

func Main() {
}
