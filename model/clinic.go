package model

import uuid "github.com/satori/go.uuid"

//Clinic saves clinic information
type Clinic struct {
	UID uuid.UUID `json:"uid"`
	//uid uuid DEFAULT uuid_generate_v4() primary key,
	AccountID uuid.UUID `json:"accountId"`
	//account_id uuid references users(uid) not null,
	Name string `json:"name"`
	//name varchar not null,
	Address string `json:"address"`
	//address varchar,
	StateID int `json:"stateId"`
	//state_id integer references state(id) not null,
	CountryID int `json:"countryId"`
	//country_id integer references country(id) not null,
	Phone string `json:"phone"`
	//phone varchar not null,
	Email string `json:"email"`
	//email varchar not null,
	Approved *bool `json:"approved"`
	//approved boolean,
	CreatedOn string `json:"createdOn"`
	//created_on timestamp,
	OnboardedOn *string `json:"onboardedOn"`
	//onboarded_on timestamp
}
