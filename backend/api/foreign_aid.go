package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type foreignAid struct {
	Country string
	Amount  float64
	Year    int
	Lat     float64
	Lng     float64
}

func (foreignAid) TableName() string {
	return "foreign_aid"
}

func (h *Handler) GetForeignAid(c *gin.Context) {
	var queryResults []foreignAid
	result := h.db.Find(&queryResults)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queryResults})
}
