package model

import uuid "github.com/satori/go.uuid"

type UserDocument struct {
	UID uuid.UUID `json:"uid" create-remove:"True" update-remove:"True"`
	//uid uuid DEFAULT uuid_generate_v4() primary key,
	UserID uuid.UUID `json:"userId" create-required:"True" update-remove:"True"`
	//user_uid uuid references users(uid) not null,
	DocumentTypeID int `json:"docTypeId" create-required:"True"`
	//doc_type_id integer not null,
	URL string `json:"url" create-required:"True"`
	//url varchar not null,
	CreatedOn string `json:"createdOn" create-remove:"True" updated-remove:"True"`
	//created_on timestamp);
}
