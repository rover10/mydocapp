package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	required := []string{"firstName", "password", "retypedPassword", "email", "phone", "genderId", "age", "countryId", "userType"}
	remove := []string{"uid", "retypedPassword", "createdOn", "updatedOn", "isActive"}

	missing := parseutil.EnsureRequired(body, required)
	body = parseutil.RemoveFields(body, remove)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"firstName", "password", "lastName", "phone", "email"}
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
	tx, err := s.DB.DB.Begin()
	fmt.Println("-->1")
	if err != nil {
		fmt.Println(err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("-->2")
	row := tx.QueryRow(query, values...)
	err = row.Scan(&user.UID, &user.FirstName, &user.Email, &user.Phone, &user.UserType, &user.Gender, &user.Country, &user.IsActive, &user.CreatedOn)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("-->3")
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		tx.Rollback()
		return context.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("-->4")
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
	tx, err := s.DB.DB.Begin()
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
		tx.Rollback()
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
	tx, err := s.DB.DB.Begin()
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
		tx.Rollback()
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
	tx, err := s.DB.DB.Begin()
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
		tx.Rollback()
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
	tx, err := s.DB.DB.Begin()
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
		tx.Rollback()
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
	floatField := []string{}
	jsonField := []string{}
	appointment := model.Appointment{}

	body, invalidType := parseutil.MapX(body, appointment, stringFields, floatField, intField, boolField, jsonField)
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
	tx, err := s.DB.DB.Begin()
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
		tx.Rollback()
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
func (s *Server) RegisterTreatment(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"appointmentId", "doctorId", "patientProblemDescription"}
	createRemove := []string{"uid", "createdOn"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringField := []string{"uid", "appointmentId", "doctorId", "patientProblemDescription", "createdOn"}
	intField := []string{""}
	floatField := []string{""}
	boolField := []string{""}
	JSONField := []string{""}
	model := model.Treatment{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "treatment")
	query = query + "RETURNING uid,appointment_id,doctor_id,patient_problem_description,created_on"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&model.UID, &model.AppointmentID, &model.DoctorID, &model.PatientProblemDesc, &model.CreatedOn)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

//RegisterDoctorReview regsiter a review by the user
func (s *Server) RegisterDoctorReview(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"appointmentId", "reviewerId", "doctorId", "rating", "review"}
	createRemove := []string{"createdOn"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringField := []string{"appointmentId", "reviewerId", "doctorId", "review", "createdOn"}
	intField := []string{""}
	floatField := []string{"rating"}
	boolField := []string{""}
	JSONField := []string{""}
	model := model.DoctorReview{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "doctor_review")
	query = query + "RETURNING appointment_id,reviewer_id,doctor_id,rating,review,created_on"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&model.AppointmentID, &model.ReviewerID, &model.DoctorID, &model.Rating, &model.Review, &model.CreatedOn)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

//UploadUserDocument upload a new document by user
func (s *Server) UploadUserDocument(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"userId", "docTypeId", "url"}
	createRemove := []string{"uid", "createdOn"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringField := []string{"uid", "userId", "url", "createdOn"}
	intField := []string{"docTypeId"}
	floatField := []string{""}
	boolField := []string{""}
	JSONField := []string{""}
	model := model.UserDocument{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "user_document")
	query = query + "RETURNING uid,user_id,doc_type_id,url,created_on"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&model.UID, &model.UserID, &model.DocumentTypeID, &model.URL, &model.CreatedOn)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

//AddTestReport register a new test report of a patient
func (s *Server) AddTestReport(context echo.Context) error {

	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"treatmentId", "docId"}
	createRemove := []string{"uid", "createdOn", "updatedOn"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringField := []string{"uid", "treatmentId", "docId", "createdOn", "updatedOn"}
	intField := []string{""}
	floatField := []string{""}
	boolField := []string{""}
	JSONField := []string{""}
	model := model.TestReport{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "test_report")
	query = query + "RETURNING uid,treatment_id,doc_id,created_on,updated_on"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	updatedOn := sql.NullString{}
	err = row.Scan(&model.UID, &model.TreatmentID, &model.DocumentID, &model.CreatedOn, &updatedOn)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

//AssignStaffRole assign a new role to a staff
func (s *Server) AssignStaffRole(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"userId", "roleId", "clinicId", "isActive"}
	createRemove := []string{"createdOn"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringField := []string{"userId", "clinicId", "createdOn"}
	intField := []string{"roleId"}
	floatField := []string{""}
	boolField := []string{"isActive"}
	JSONField := []string{""}
	model := model.StaffRole{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "staff_role")
	query = query + "RETURNING user_id,role_id,clinic_id,created_on,is_active"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&model.UserID, &model.RoleID, &model.ClinicID, &model.CreatedOn, &model.IsActive)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

//AddDoctorQualification add doctor qualification
func (s *Server) AddDoctorQualification(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}

	createRequired := []string{"userId", "qualificationId", "certificateDoc", "university", "universityId"}
	createRemove := []string{"verified"}
	body = parseutil.RemoveFields(body, createRemove)
	missing := parseutil.EnsureRequired(body, createRequired)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	// updateRequired := []string{""}
	// updateRemove := []string{"uuid.UUID", "int", "string", "string", "bool"}
	// body = parseutil.RemoveFields(body, updateRemove)
	// missing := parseutil.EnsureRequired(body, updateRequired)
	// if len(missing) != 0 {
	// 	log.Println("missing", missing)
	// 	return context.JSON(http.StatusBadRequest, missing)
	// }

	stringField := []string{"userId", "createdOn", "certificateDoc"}
	intField := []string{"qualificationId"}
	floatField := []string{""}
	boolField := []string{"verified"}
	JSONField := []string{""}
	model := model.DoctorQualification{}

	body, invalidType := parseutil.MapX(body, model, stringField, floatField, intField, boolField, JSONField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	query, values := querybuilder.BuildInsertQuery(body, "doctor_qualification")
	query = query + "RETURNING user_id,qualification_id,created_on,certificate_doc,verified"

	tx, err := s.DB.DB.Begin()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	row := tx.QueryRow(query, values...)

	err = row.Scan(&model.UserID, &model.QualificationID, &model.CreatedOn, &model.CertificateDoc, &model.Verified)

	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("\nDatabase Commit Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, model)
}

func (s *Server) Login(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	username := body["username"] //context.FormValue("username")
	password := body["password"] //context.FormValue("password")
	// Check in your db if the user exists or not
	email := username //context.Get("username")
	emailStr, ok := email.(string)
	fmt.Println("->>" + emailStr)
	fmt.Println(email)
	fmt.Println(context)

	if ok && emailStr != "" {
		fmt.Println("->> ))" + emailStr)
		user, err := s.DB.RetriveUserCred(emailStr)
		if err == sql.ErrNoRows {
			return errors.New("User does not exit")
		} else if err != nil {
			return errors.New("Server error")
		}
		fmt.Println("->> 2))" + emailStr)
		if username == user.Email && password == user.Password {
			fmt.Println("->> )) 3" + emailStr)
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)
			// Set claims
			// This is the information which frontend can use
			// The backend can also decode the token and get admin etc.
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = user.FirstName
			claims["uid"] = user.UID
			claims["name"] = user.Email
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			// Generate encoded token and send it as response.
			// The signing string should be secret (a generated UUID          works too)

			t, err := token.SignedString([]byte(SECRET_PASSWORD))
			if err != nil {
				return err
			}
			fmt.Println("->> )) 3" + emailStr)
			// Generate encoded refrest token and send it as response.
			refreshToken := jwt.New(jwt.SigningMethodHS256)
			rtClaims := refreshToken.Claims.(jwt.MapClaims)
			rtClaims["sub"] = 1
			rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
			rt, err := refreshToken.SignedString([]byte(SECRET_PASSWORD))
			if err != nil {
				return err
			}

			return context.JSON(http.StatusOK, map[string]string{
				"token":         t,
				"refresh_token": rt,
			})
		}
	}

	return echo.ErrUnauthorized
}
