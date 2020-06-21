package model

import uuid "github.com/satori/go.uuid"

//Patient register patient detail
type Patient struct {
	UID uuid.UUID `json:"uid"`
	//uid uuid DEFAULT uuid_generate_v4() primary key,
	AccountID uuid.UUID `json:"accountId"`
	//account_id uuid references users(uid),
	FirstName string `json:"firstName"`
	//first_name varchar not null,
	LastName string `json:"lastName"`
	//last_name varchar,
	GenderID int `json:"genderId"`

	CountryID int `json:"countryId"`
	//gender varchar,
	Age int `json:"age"`
	//age int,
	Phone string `json:"phone"`
	//phone varchar not null,
	Email string `json:"email"`
	//email varchar not null,
	AnyExistingMedicalCondition MedicalCondition `json:"anyExistingMedicalCondition"`
	//existing_condition jsonb,
	CreatedOn string `json:"createdOn"`
	//created_on timestamp default now(),
	UpdatedOn *string `json:"updatedOn"`
	//updated_on timestamp
}

//MedicalCondition store medical contions
type MedicalCondition struct {
	Conditions *[]Condition `json:"conditions"`
}

//Condition
type Condition struct {
	Name        string  `json:"name"`
	ID          int     `json:"id"`
	Description *string `json:"description"`
	From        *int    `json:"from"`
}
