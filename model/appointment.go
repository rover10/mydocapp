package model

import uuid "github.com/satori/go.uuid"

type Appointment struct {
	UID uuid.UUID `json:"uid"`
	//uid uuid DEFAULT uuid_generate_v4() primary key,
	AccountID uuid.UUID `json:"accountId"`
	ClinicID  uuid.UUID `json:"clinicId"`
	//clinic_id uuid references clinic(uid) not null,
	//account_id uuid references users(uid) not null,
	//Doctor can be assigned later so it is nullable
	DoctorID *uuid.UUID `json:"doctorId"`
	//doctor_id uuid references users(uid),
	PatientID uuid.UUID `json:"patientId"`
	DiseaseID *int      `json:"diseaseId"`
	//disease_id integer references disease(id),
	//booking_date timestamp default now(),
	SlotDateTime string `json:"slotDateTime"`
	//slot_date_time timestamp not null,
	ContactPhone string `json:"contactPhone"`
	//contact_phone varchar);
	NoShow bool `json:"noShow"`
	//no_show boolean default false not null,
	CreatedOn string `json:"createdOn"`
	UpdatedOn string `json:"-"`
	Cancelled bool   `json:"cancelled"`
}
