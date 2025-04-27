package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var renaming = map[string]string{
	"Department of the Treasury":                  "Treasury",
	"Department of Health and Human Services":     "Health and Human",
	"Department of Defense":                       "Defense",
	"Social Security Administration":              "Social Security",
	"Department of Veterans Affairs":              "Veterans Affairs",
	"Department of Agriculture":                   "Agriculture",
	"Office of Personnel Management":              "OPM",
	"Department of Housing and Urban Development": "Housing",
	"Department of Transportation":                "Transportation",
	"Department of Homeland Security":             "Homeland Security",
	"Department of Energy":                        "Energy",
	"Department of Commerce":                      "Commerce",
	"Department of Education":                     "Education",
	"Environmental Protection Agency":             "Environmental",
	"Department of the Interior":                  "Interior",
	"Department of State":                         "State",
	"General Services Administration":             "General Services",
	"Department of Justice":                       "Justice",
	"Department of Labor":                         "Labor",
	"Pension Benefit Guaranty Corporation":        "Pension",
}

var colors = []string{
	"0, 128, 128",  // Teal
	"255, 99, 71",  // Tomato
	"124, 252, 0",  // Lawn Green
	"70, 130, 180", // Steel Blue
	"255, 215, 0",  // Gold
	"0, 191, 255",  // Deep Sky Blue
	"255, 69, 0",   // Orange Red
	"138, 43, 226", // Blue Violet
	"60, 179, 113", // Medium Sea Green
	"218, 165, 32", // Golden Rod
}

func getNameAndColor(i int, agency string) (string, string) {
	color := colors[i%9]
	if rename, exists := renaming[agency]; exists {
		return rename, color
	}
	return agency, color
}

type AgencyBudget struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Agency string  `json:"agency"`
	Budget float64 `json:"budget"`
}

func (AgencyBudget) TableName() string {
	return "agency_budget"
}

type pieResult struct {
	Proportion float64 `json:"proportion"`
	Amount     float64 `json:"amount"`
	Agency     string  `json:"agency"`
	Color      string  `json:"color"`
}

type returnResult struct {
	FirstPie  []pieResult `json:"first_pie"`
	SecondPie []pieResult `json:"second_pie"`
	Table     []pieResult `json:"table"`
}

func addEntry(name, color string, remainingBudget float64, budget float64) pieResult {
	return pieResult{
		Proportion: budget / remainingBudget,
		Amount:     budget,
		Agency:     name,
		Color:      color,
	}

}

func (h *Handler) GetAgencyData(c *gin.Context) {
	var agencyBudgets []AgencyBudget
	result := h.db.Order("budget DESC").Find(&agencyBudgets)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var totalBudget float64
	var mainRemaingBudget float64
	var otherRmaingBudget float64
	for i, agency := range agencyBudgets {
		totalBudget += agency.Budget
		if i > 8 {
			mainRemaingBudget += agency.Budget
		}
		if i > 17 {
			otherRmaingBudget += agency.Budget
		}
	}

	firstPie := []pieResult{}
	secondPie := []pieResult{}
	table := []pieResult{}
	var color string
	var name string

	for i, agency := range agencyBudgets {
		name, color = getNameAndColor(i, agency.Agency)
		if i < 9 {
			firstPie = append(
				firstPie,
				addEntry(name, color, totalBudget, agency.Budget),
			)
		} else if i < 18 {
			secondPie = append(
				secondPie,
				addEntry(name, color, totalBudget, agency.Budget),
			)
		} else {
			table = append(
				table,
				addEntry(name, color, totalBudget, agency.Budget),
			)
		}
	}

	name = fmt.Sprintf("%d other agencies", len(agencyBudgets)-9)
	firstPie = append(
		firstPie,
		addEntry(name, colors[9], totalBudget, mainRemaingBudget),
	)
	name = fmt.Sprintf("%d other agencies", len(agencyBudgets)-18)
	secondPie = append(
		secondPie,
		addEntry(name, colors[9], totalBudget, otherRmaingBudget),
	)

	c.JSON(http.StatusOK, returnResult{
		FirstPie:  firstPie,
		SecondPie: secondPie,
		Table:     table,
	})
}
