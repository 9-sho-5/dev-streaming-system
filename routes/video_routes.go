package routes

import (
	"dev-streaming-system/handlers/video"

	"github.com/gin-gonic/gin"
)

func RegisterVideoRoutes(router *gin.Engine) {
	// ビデオ表示画面
	router.GET("/video", video.VideoHandler)
}
