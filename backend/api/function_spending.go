package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type functionSpending struct {
	Year   int     `json:"year"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func (functionSpending) TableName() string {
	return "function_spending"
}

func (h *Handler) GetFunctionSpending(c *gin.Context) {
	var infos []functionSpending
	returnData := make(map[string][]functionSpending)
	result := h.db.Order("year").Find(&infos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for _, item := range infos {
		returnData[item.Name] = append(
			returnData[item.Name],
			item,
		)
	}
	c.JSON(http.StatusOK, gin.H{"data": returnData})
}
