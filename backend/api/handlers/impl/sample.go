package impl

import (
	"net/http"

	openapi "api/controllers/restapi"
	"api/handlers"

	"github.com/gin-gonic/gin"
)

type sampleHandlersImpl struct {
}

var _ handlers.SampleHandlers = &sampleHandlersImpl{}

func NewSampleHandlers() handlers.SampleHandlers {
	return &sampleHandlersImpl{}
}

func (h *sampleHandlersImpl) PostSample(c *gin.Context) {
	req := &openapi.AuthRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
