package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthenticationHandlers interface {
	PostLogin(c *gin.Context)
	PostRefreshToken(c *gin.Context)
	PostLogout(c *gin.Context)
}
