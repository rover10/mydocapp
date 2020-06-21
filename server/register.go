package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rover10/mydocapp.git/model"
	"github.com/rover10/mydocapp.git/parseutil"
	"github.com/rover10/mydocapp.git/querybuilder"
)

// Ping - This function will ping the echo server
func (s *Server) Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Health": "OK"})
}

// RegisterUser register a new user
func (s *Server) RegisterUser(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{"firstName", "email", "phone", "genderId", "age", "countryId", "userType"}
	remove := []string{"uid", "createdOn", "updatedOn", "isActive"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"firstName", "lastName", "phone", "email"}
	intFields := []string{"userType", "genderId", "countryId"}
	user := model.User{}
	body, invalidType := parseutil.MapX(body, user, stringFields, nil, intFields, nil, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "users")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING uid, first_name, email, phone, user_type, gender_id, country_id, is_active, created_on"

	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)
	err = row.Scan(&user.UID, &user.FirstName, &user.Email, &user.Phone, &user.UserType, &user.Gender, &user.Country, &user.IsActive, &user.CreatedOn)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, user)
}

//RegisterPatient register a new patient
func (s *Server) RegisterPatient(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	// Move to a normalize table to make modification and reitreval easy instead of JSON
	/*
			"anyExistingMedicalCondition": {
			"conditions": [
					{
						"name": "diabetes",
						"id": 1,
						"description": "Unknown",
						"from": "2018-06-21T22:28:12.608205Z"
					}
				]
		    }
	*/
	fmt.Println(body)
	required := []string{"accountId", "firstName", "email", "phone", "genderId", "age", "countryId"}
	remove := []string{"uid", "createdOn", "updatedOn"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"accountId", "firstName", "lastName", "phone", "email"}
	intFields := []string{"genderId", "age", "countryId"}
	//jsonFields := []string{"anyExistingMedicalCondition"}
	patient := model.Patient{}
	body, invalidType := parseutil.MapX(body, patient, stringFields, nil, intFields, nil, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "patient")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING uid, account_id, first_name, last_name, age, email, phone, gender_id, country_id, created_on"

	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)
	lastName := sql.NullString{}

	err = row.Scan(&patient.UID, &patient.AccountID, &patient.FirstName, &lastName, &patient.Age, &patient.Email, &patient.Phone, &patient.GenderID, &patient.CountryID, &patient.CreatedOn)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	if lastName.Valid {
		patient.LastName = lastName.String
	}
	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, patient)
}

//RegisterDoctor register a new doctor
func (s *Server) RegisterDoctor(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{"accountId"}
	remove := []string{"rating", "approved", "onboardedOn"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"accountId", "practiceStartDate"}
	floatFields := []string{"fee"}
	doctor := model.Doctor{}
	body, invalidType := parseutil.MapX(body, doctor, stringFields, floatFields, nil, nil, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "doctor")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING account_id, fee, practice_start_date"

	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)
	fee := sql.NullFloat64{}
	practiceStartDate := sql.NullString{}
	err = row.Scan(&doctor.AccountID, &fee, &practiceStartDate)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	if fee.Valid {
		doctor.Fee = &fee.Float64
	}
	if practiceStartDate.Valid {
		doctor.PracticeStartDate = &practiceStartDate.String
	}

	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, doctor)

}

//RegisterClinic register a new clinic
func (s *Server) RegisterClinic(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{"accountId", "name", "address", "stateId", "countryId", "phone", "email"}
	remove := []string{"uid", "approved", "createdOn", "onboardedOn"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"accountId", "name", "address", "phone", "email"}
	intFields := []string{"stateId", "countryId"}
	clinic := model.Clinic{}
	body, invalidType := parseutil.MapX(body, clinic, stringFields, nil, intFields, nil, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "clinic")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING uid, account_id, name,  address, state_id, country_id, phone, email, created_on"
	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&clinic.UID, &clinic.AccountID, &clinic.Name, &clinic.Address, &clinic.StateID, &clinic.CountryID, &clinic.Phone, &clinic.Email, &clinic.CreatedOn)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}

	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, clinic)
}

//RegisterStaff register new staff of clinics
func (s *Server) RegisterStaff(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{"accountId", "clinicId"}
	remove := []string{"createdOn", "isActive"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"accountId", "clinicId", "createdOn"}
	boolField := []string{"isActive"}
	staff := model.Staff{}
	body, invalidType := parseutil.MapX(body, staff, stringFields, nil, nil, boolField, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "staff")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING account_id, clinic_id, created_on, is_active"
	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&staff.AccountID, &staff.ClinicID, &staff.CreatedOn, &staff.IsActive)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}

	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, staff)
}

//BookAppointment book a new appointment for consultation
func (s *Server) BookAppointment(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{"accountId", "clinicId", "patientId", "slotDateTime", "contactPhone"}
	remove := []string{"uid", "createdOn", "noShow"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"accountId", "clinicId", "patientId", "slotDateTime", "contactPhone"}
	boolField := []string{"noShow"}
	intField := []string{"diseaseId"}
	appointment := model.Appointment{}

	body, invalidType := parseutil.MapX(body, appointment, stringFields, nil, intField, boolField, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "appointment")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING uid, account_id, clinic_id, patient_id, disease_id, slot_date_time, contact_phone, no_show, created_on"
	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	tx, err := s.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)
	diseaseID := sql.NullInt64{}
	err = row.Scan(&appointment.UID, &appointment.AccountID, &appointment.ClinicID, &appointment.PatientID, &diseaseID, &appointment.SlotDateTime, &appointment.ContactPhone, &appointment.NoShow, &appointment.CreatedOn)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	if diseaseID.Valid {
		dval := diseaseID.Int64
		intdVal := int(dval)
		appointment.DiseaseID = &intdVal
	}
	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, appointment)

}

//RegisterTreatment register a new problem_description, diagnosis, prescription and tests
func RegisterTreatment(context echo.Context) error {
	return nil
}

//RegisterDoctorReview regsiter a review by the user
func RegisterDoctorReview(context echo.Context) error {
	return nil
}

//UploadUserDocument upload a new document by user
func UploadUserDocument(context echo.Context) error {
	return nil
}

//AddTestReport register a new test report of a patient
func AddTestReport(context echo.Context) error {
	return nil
}

//AssignStaffRole assign a new role to a staff
func AssignStaffRole(context echo.Context) error {
	return nil
}

//AddDoctorQualification add doctor qualification
func AddDoctorQualification(context echo.Context) error {
	return nil
}
