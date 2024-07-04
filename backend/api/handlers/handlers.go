package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthenticationAPIHandlers interface {
	PostLogin(c *gin.Context)
	PostLogout(c *gin.Context)
	PostRefreshToken(c *gin.Context)
}
