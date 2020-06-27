package model

import uuid "github.com/satori/go.uuid"

type TestReport struct {
	UID uuid.UUID `json:"uid" create-remove:"True" update-remove:"True"`
	//uid uuid DEFAULT uuid_generate_v4() primary key,
	TreatmentID uuid.UUID `json:"treatmentId" create-required:"True"`
	//appointment_id uuid references appointment(uid),
	DocumentID uuid.UUID `json:"docId" create-required:"True"`
	//doc_id uuid references user_document(uid));
	CreatedOn string `json:"createdOn" create-remove:"True" update-remove:"True"`
	//created_on timestamp default now(),
	UpdatedOn *string `json:"updatedOn" create-remove:"True" update-remove:"True"`
	//updated_on timestamp,
}
