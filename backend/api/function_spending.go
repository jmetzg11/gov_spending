package api

import (
	"net/http"
	"sort"

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

type agencySpending struct {
	Year   int     `json:"year"`
	Amount float64 `json:"amount"`
}

func (h *Handler) GetFunctionSpending(c *gin.Context) {
	var infos []functionSpending
	result := h.db.Order("year").Find(&infos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	agencies := make(map[string]struct{})
	years := make(map[int]struct{})
	data := make(map[string][]agencySpending)

	for _, item := range infos {
		agencies[item.Name] = struct{}{}
		years[item.Year] = struct{}{}
		data[item.Name] = append(data[item.Name], agencySpending{
			Year:   item.Year,
			Amount: item.Amount / 1e9,
		})
	}

	for key := range data {
		sort.Slice(data[key], func(i, j int) bool {
			return data[key][i].Year < data[key][j].Year
		})
	}

	agencyList := make([]string, 0, len(agencies))
	for agency := range agencies {
		agencyList = append(agencyList, agency)
	}
	sort.Slice(agencyList, func(i, j int) bool {
		// Get the most recent year for each agency
		agencyI := agencyList[i]
		agencyJ := agencyList[j]

		// Find the most recent data point for each
		latestYearI := data[agencyI][len(data[agencyI])-1]
		latestYearJ := data[agencyJ][len(data[agencyJ])-1]

		// Sort by amount in descending order
		return latestYearI.Amount > latestYearJ.Amount
	})

	yearList := make([]int, 0, len(years))
	for year := range years {
		yearList = append(yearList, year)
	}
	sort.Ints(yearList)

	c.JSON(http.StatusOK, gin.H{
		"data":     data,
		"years":    yearList,
		"agencies": agencyList,
	})
}
