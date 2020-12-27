package response

import uuid "github.com/satori/go.uuid"

type Appointment struct {
	UID          uuid.UUID  `json:"uid"`
	AccountID    uuid.UUID  `json:"-"`            // Booked by
	ClinicID     uuid.UUID  `json:"-"`            // Booked Clinic
	DoctorID     *uuid.UUID `json:"-"`            // Doctor chose
	PatientID    uuid.UUID  `json:"-"`            // Booked for
	DiseaseID    *int       `json:"diseaseId"`    // Booked for disease
	SlotDateTime string     `json:"slotDateTime"` // Booking date
	ContactPhone string     `json:"contactPhone"`
	NoShow       bool       `json:"noShow"`
	CreatedOn    string     `json:"createdOn"`
	Cancelled    bool       `json:"cancelled"`
	Clinic       Clinic     `json:"clinic" gorm:"foreignKey:ClinicID; references:UID"` // references is the key in this model Clinic being referred by the Fk ClinicID in this  model
	Doctor       *Doctor    `json:"doctor" gorm:"foreignKey:DoctorID; references:AccountID"`
	Patient      Patient    `json:"patient" gorm:"foreignKey:PatientID; references:UID"`
}

// User
type User struct {
	UID       uuid.UUID `json:"uid" gorm:"PRIMARY_KEY"`
	FirstName string    `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Gender    int       `json:"genderId"`
}

// Clinic
type Clinic struct {
	UID       uuid.UUID `json:"uid" gorm:"PRIMARY_KEY"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CountryID int       `json:"countryId"`
	StateID   int       `json:"stateId"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
}

// Doctor
type Doctor struct {
	AccountID uuid.UUID `json:"accountId" gorm:"PRIMARY_KEY"`
	Rating    *float64  `json:"rating"`
}

// Disease
type Disease struct{}

// Patient
type Patient struct {
	UID       uuid.UUID `json:"uid" gorm:"PRIMARY_KEY"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
}
