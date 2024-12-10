package main

import (
	"dev-streaming-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルータを作成
	router := gin.Default()

	// ルートにアクセスしたときに "Hello! World" を表示
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello! World",
		})
	})

	// ルーティングを登録
	routes.RegisterVideoRoutes(router)

	// サーバー起動
	router.Run(":8080")
}
