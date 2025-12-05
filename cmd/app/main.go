package main

import (
	"github.com/gin-gonic/gin"
	"2512-hack-for-change-backend/internal/handler"
)

func main() {
	r := gin.Default()

	// 動作確認用のAPI
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 足し算用のHandlerを初期化
	addHandler := handler.NewAddHandler()
	// GET /add?a=1&b=2 にアクセスがきたら addHandler.Addを実行する
	r.GET("/add", addHandler.Add)

	r.Run() // デフォルトは :8080
}
