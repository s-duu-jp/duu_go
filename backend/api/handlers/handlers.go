package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthenticationHandlers interface {
	PostLogin(c *gin.Context)
	PostLogout(c *gin.Context)
}

type SampleHandlers interface {
	PostSample(c *gin.Context)
}
