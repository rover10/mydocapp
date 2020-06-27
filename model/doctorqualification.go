package model

import uuid "github.com/satori/go.uuid"

type DoctorQualification struct {
	UserID uuid.UUID `json:"userId" create-required:"True" update-remove:"True"`
	//user_id uuid references users(uid) not null,
	QualificationID int `json:"qualificationId" create-required:"True" update-remove:"True"`
	//qualification_id integer references qualification(id),
	CreatedOn string `json:"createdOn" create-removed:"True" update-remove:"True"`
	//created_on timestamp default now() not null,
	CertificateDoc string `json:"certificateDoc" create-required:"True" update-remove:"True"`
	//certificate_doc uuid references user_document(uid),
	Verified bool `json:"verified" create-remove:"True" update-remove:"True"`
	//verified boolean not null default false)
}
