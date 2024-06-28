package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	AuthenticationHandlers
	SampleHandlers
}

type AuthenticationHandlers interface {
	PostLogin(c *gin.Context)
	PostRefreshToken(c *gin.Context)
	PostLogout(c *gin.Context)
}

type SampleHandlers interface {
	PostSample(c *gin.Context)
}
