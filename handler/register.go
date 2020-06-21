package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rover10/mydocapp.git/model"
	"github.com/rover10/mydocapp.git/parseutil"
	"github.com/rover10/mydocapp.git/querybuilder"
)

// Ping - This function will ping the echo server
func Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Health": "OK"})
}

// RegisterUser register a new user
func RegisterUser(context echo.Context) error {
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	//
	required := []string{"firstName", "email", "phone", "gender", "age", "country", "userType"}
	remove := []string{"uid", "createdOn", "updatedOn", "isActive"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"firstName", "lastName", "phone", "email"}
	intFields := []string{"userType", "gender", "country"}
	user := model.User{}
	body, invalidType := parseutil.MapX(body, user, stringFields, nil, intFields, nil)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "users")
	fmt.Println(query)
	fmt.Println(values)
	// Execute query
	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, body)
}

//RegisterPatient register a new patient
func RegisterPatient(context echo.Context) error {
	return nil
}

//RegisterDoctor register a new doctor
func RegisterDoctor(context echo.Context) error {
	return nil
}

//RegisterClinic register a new clinic
func RegisterClinic(context echo.Context) error {
	return nil
}

//RegisterStaff register new staff of clinics
func RegisterStaff(context echo.Context) error {
	return nil
}

//BookAppointment book a new appointment for consultation
func BookAppointment(context echo.Context) error {
	return nil

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
