package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetForeignAid(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "working"})
}
