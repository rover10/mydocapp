package model

import uuid "github.com/satori/go.uuid"

type StaffRole struct {
	UserID uuid.UUID `json:"userId" create-required:"True" update-remove:"True"`
	//user_id uuid references users(uid) not null,
	RoleID int `json:"roleId" create-required:"True" update-remove:"True"`
	//role_id integer references roles(id) not null,
	ClinicID  uuid.UUID `json:"clinicId" create-required:"True" update-remove:"True"`
	CreatedOn string    `json:"createdOn" create-remove:"True" update-remove:"True"`
	//created_on timestamp default now() not null,
	IsActive bool `json:"isActive" create-required:"True"`
	//is_active boolean not null default true);
}
