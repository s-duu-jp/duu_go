package impl

import (
	"net/http"

	openapi "api/controllers/restapi"

	"github.com/gin-gonic/gin"
)

func (h *handlersImpl) PostSample(c *gin.Context) {
	req := &openapi.PostLoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
