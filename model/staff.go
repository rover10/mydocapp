package model

import uuid "github.com/satori/go.uuid"

type Staff struct {
	AccountID uuid.UUID `json:"accountId"`
	//account_id uuid references users(uid),
	ClinicID uuid.UUID `json:"clinicId"`
	//clinic_id uuid references clinic(uid),
	CreatedOn string `json:"createdOn"`
	//created_on timestamp,
	IsActive bool `json:"isActive"`
	//is_active boolean default true
}
