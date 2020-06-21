package model

import uuid "github.com/satori/go.uuid"

type User struct {
	UID       uuid.UUID `json:"uid"`
	UserType  int       `json:"userType"`
	IsActive  *bool     `json:"isActive"`
	FirstName string    `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Gender    int       `json:"genderId"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Country   *int      `json:"countryId"`
	CreatedOn string    `json:"createdOn"`
	UpdatedOn *string   `json:"updatedOn"`
}
