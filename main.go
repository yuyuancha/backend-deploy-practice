package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("設定 ENV 發生錯誤:", err.Error())
	}
	fmt.Println("APP_ENV:", os.Getenv("APP_ENV"))
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
