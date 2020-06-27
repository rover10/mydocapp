package model

import uuid "github.com/satori/go.uuid"

type Treatment struct {
	UID uuid.UUID `json:"uid" create-remove:"True" update-remove:"True"`
	// uid uuid DEFAULT uuid_generate_v4() primary key,
	AppointmentID uuid.UUID `json:"appointmentId" create-required:"True"`
	//appointment_id uuid references appointment(uid) not null,
	DoctorID uuid.UUID `json:"doctorId" create-required:"True" update-remove:"True"`
	//doctor_id uuid references users(uid) not null,
	PatientProblemDesc string `json:"patientProblemDescription" create-required:"True" update-remove:"True"`
	//patient_problem_description varchar);
	CreatedOn string `json:"createdOn" create-remove:"True"`
}
