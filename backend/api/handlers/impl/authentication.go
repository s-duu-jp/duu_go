package impl

import (
	"net/http"

	openapi "api/controllers/restapi"

	"github.com/gin-gonic/gin"
)

// ログイン
func (h *handlersImpl) PostLogin(c *gin.Context) {
	req := &openapi.AuthRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// トークンの再発行
func (h *handlersImpl) PostRefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ログアウト
func (h *handlersImpl) PostLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
