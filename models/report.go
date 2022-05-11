package models

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	ScammerPhoneNumber string `json:"scammerphonenumber"`
	Description        string `json:"description"`
	Category           string `json:"category"`
	Author             string `json:"author"`
	ReportDate         string `json:"reportDate"`
	NumberOfReports    int    `json:"numberOfReports"`
}
