package model

import uuid "github.com/satori/go.uuid"

type DoctorReview struct {
	AppointmentID uuid.UUID `json:"appointmentId" create-required:"True" update-remove:"True"`
	//appointment_id uuid references appointment(uid),
	ReviewerID uuid.UUID `json:"reviewerId" create-required:"True" update-remove:"True"`
	// reviewer_id uuid references users(uid),
	DoctorID uuid.UUID `json:"doctorId" create-required:"True" update-remove:"True"`
	// doctor_id uuid references users(uid),
	Rating float64 `json:"rating" create-required:"True" update-remove:"True"`
	// rating float check (rating >= 0) check (rating <= 5),
	Review string `json:"review" create-required:"True" update-remove:"True"`
	// review varchar, review_date timestamp);
	CreatedOn string `json:"createdOn" create-remove:"True" update-remove:"True"`
}
