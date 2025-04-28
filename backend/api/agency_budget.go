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
	"Agency for International Development":        "International Development",
}

var colors = []string{
	"45, 49, 66",   // Dark Blue-Gray
	"191, 52, 52",  // Deep Red
	"52, 138, 167", // Teal Blue
	"147, 76, 123", // Plum
	"45, 92, 126",  // Navy Blue
	"178, 84, 34",  // Burnt Orange
	"91, 43, 115",  // Deep Purple
	"52, 125, 87",  // Forest Green
	"128, 101, 49", // Olive Brown
	"76, 82, 96",   // Slate Gray
}

func getNameAndColor(agency string) string {
	if rename, exists := renaming[agency]; exists {
		return rename
	}
	return agency
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

func addEntry(name string, i int, remainingBudget float64, budget float64) pieResult {
	color := colors[i%10]
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
	var name string

	for i, agency := range agencyBudgets {
		name = getNameAndColor(agency.Agency)
		if i < 9 {
			firstPie = append(
				firstPie,
				addEntry(name, i, totalBudget, agency.Budget),
			)
		} else if i < 18 {
			secondPie = append(
				secondPie,
				addEntry(name, i-9, totalBudget, agency.Budget),
			)
		} else {
			table = append(
				table,
				addEntry(name, i, totalBudget, agency.Budget),
			)
		}
	}

	name = fmt.Sprintf("%d other agencies", len(agencyBudgets)-9)
	firstPie = append(
		firstPie,
		addEntry(name, 9, totalBudget, mainRemaingBudget),
	)
	name = fmt.Sprintf("%d other agencies", len(agencyBudgets)-18)
	secondPie = append(
		secondPie,
		addEntry(name, 9, totalBudget, otherRmaingBudget),
	)

	c.JSON(http.StatusOK, returnResult{
		FirstPie:  firstPie,
		SecondPie: secondPie,
		Table:     table,
	})
}
