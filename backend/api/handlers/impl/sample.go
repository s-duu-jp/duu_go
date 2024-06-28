package impl

import (
	"net/http"

	openapi "api/controllers/restapi"
	"api/handlers"

	"github.com/gin-gonic/gin"
)

type sampleHandlers struct {
}

var _ handlers.SampleHandlers = &sampleHandlers{}

func NewSampleHandlers() *sampleHandlers {
	return &sampleHandlers{}
}

// ログイン
func (h *sampleHandlers) PostSample(c *gin.Context) {
	req := &openapi.PostLoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
