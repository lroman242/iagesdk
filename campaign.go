package iagesdk

import "time"

// Campaign describe iAGE`s campaign data structure
type Campaign struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Status      int       `json:"status"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Budget      float64   `json:"budget,omitempty"`
	DailyBudget float64   `json:"dailyBudget,omitempty"`
	BudgetEven  bool      `json:"budgetEven,omitempty"`
}
