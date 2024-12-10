package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VideoHandler はリクエストに応じて動画を配信します
func VideoHandler(c *gin.Context) {
	// ユーザが存在する場合は200を返す
	c.JSON(http.StatusOK, gin.H{
		"message": "Video page accessed successfully",
	})
}
