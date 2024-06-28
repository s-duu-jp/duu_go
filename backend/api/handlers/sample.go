package handlers

import (
	"github.com/gin-gonic/gin"
)

type SampleHandlers interface {
	PostSample(c *gin.Context)
}
