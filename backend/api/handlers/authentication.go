package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ログイン
func PostLogin(c *gin.Context) {
	req := factory()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// トークンの再発行
func PostRefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ログアウト
func PostLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
