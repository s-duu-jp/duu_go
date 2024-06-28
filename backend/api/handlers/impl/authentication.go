package impl

import (
	"net/http"

	openapi "api/controllers/restapi"
	"api/handlers"

	"github.com/gin-gonic/gin"
)

type authenticationHandlers struct {
}

var _ handlers.AuthenticationHandlers = &authenticationHandlers{}

func NewAuthenticationHandlers() *authenticationHandlers {
	return &authenticationHandlers{}
}

// ログイン
func (h *authenticationHandlers) PostLogin(c *gin.Context) {
	req := &openapi.PostLoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// トークンの再発行
func (h *authenticationHandlers) PostRefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ログアウト
func (h *authenticationHandlers) PostLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
