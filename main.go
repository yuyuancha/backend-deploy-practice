package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	run()
}

// Run 啟動伺服器服務
func run() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run()
	if err != nil {
		log.Fatalln("開啟 Gin 服務失敗:", err.Error())
	}
}
