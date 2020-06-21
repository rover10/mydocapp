package model

import uuid "github.com/satori/go.uuid"

//Doctor create doctor
type Doctor struct {
	AccountID         uuid.UUID `json:"accountId"`
	Fee               *float64  `json:"fee"`
	Rating            *float64  `json:"rating"`
	PracticeStartDate *string   `json:"practiceStartDate"`
	Approved          *bool     `json:"approved"`
	OnboardedOn       *string   `json:"onboardedOn"`
}
