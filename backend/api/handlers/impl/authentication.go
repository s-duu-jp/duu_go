package impl

import (
	openapi "api/controllers/restapi"

	"api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authenticationHandlersImpl struct {
}

var _ handlers.AuthenticationAPIHandlers = &authenticationHandlersImpl{}

func NewAuthenticationHandlers() handlers.AuthenticationAPIHandlers {
	return &authenticationHandlersImpl{}
}

// ログイン
func (h *authenticationHandlersImpl) PostLogin(c *gin.Context) {
	req := &openapi.AuthRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// トークンの再発行
func (h *authenticationHandlersImpl) PostRefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ログアウト
func (h *authenticationHandlersImpl) PostLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
