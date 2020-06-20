package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping - This function will ping the echo server
func Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Health": "OK"})
}

// RegisterUser register a new user
func RegisterUser(context echo.Context) error {
	return nil
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
