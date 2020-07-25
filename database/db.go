package database

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rover10/mydocapp.git/model"
)

type DocDB struct {
	DBORM *gorm.DB
	DB    *sql.DB
}

func (db *DocDB) RetriveUserCred(email string) (model.User, error) {
	query := " SELECT uid, password, first_name, email, phone, user_type, country_id, is_active from users where email = $1"
	fmt.Println(query)
	//rr := db.DB.Begin().Exec(query, email)
	//db.DB.Raw()
	row := db.DBORM.Raw(query, email)

	user := model.User{}
	row.Scan(&user)
	// err := row.Scan(&user.UID, &user.Password, &user.FirstName, &user.Email, &user.Phone, &user.UserType, &user.Country, &user.IsActive)
	// if err != nil {
	// 	log.Printf("\nDatabase Error: %+v", err)
	// 	return user, err
	// }
	return user, nil
}

func (db *DocDB) RetrieveAppointment(uid string) (model.Appointment, error) {
	query := " SELECT uid, account_id, clinic_id, patient_id, disease_id, slot_date_time, contact_phone, no_show, created_on from appointment WHERE uid = $1"
	//diseaseID := sql.NullInt64{}
	row := db.DBORM.Raw(query, uid)
	appointment := model.Appointment{}
	row.Scan(&appointment)
	//err := row.Scan(&appointment.UID, &appointment.AccountID, &appointment.ClinicID, &appointment.PatientID, &diseaseID, &appointment.SlotDateTime, &appointment.ContactPhone, &appointment.NoShow, &appointment.CreatedOn)
	// if err != nil {
	// 	log.Printf("\nDatabase Error: %+v", err)
	// 	return appointment, err
	// }
	// if diseaseID.Valid {
	// 	val := int(diseaseID.Int64)
	// 	appointment.DiseaseID = &val
	// }
	return appointment, nil
}

func (db *DocDB) RetrieveDoctorDetail(sid string) error {
	return nil
}

func (db *DocDB) RetrieveTreatment(uid string) error {
	return nil
}

func (db *DocDB) RetrieveTestResult(uid string) error {
	return nil
}

func (db *DocDB) RetrieveAllAppointment(uid string) ([]model.Appointment, error) {
	var allAppointments []model.Appointment
	return allAppointments, nil
}
