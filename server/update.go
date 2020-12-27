package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rover10/model"
	"github.com/rover10/parseutil"
	"github.com/rover10/querybuilder"
	"github.com/rover10/token"
)

//UpdateDoctor update
func UpdateDoctor(context echo.Context) error {
	return nil
}

//UpdateUser update user
func UpdateUser(context echo.Context) error {
	return nil
}

//UpdatePatient update patient info
func UpdatePatient(context echo.Context) error {
	return nil
}

//UpdateClinic update clinic
func UpdateClinic(context echo.Context) error {
	return nil
}

//UpdateStaff update staff
func UpdateStaff(context echo.Context) error {
	return nil
}

//UpdateAppointment user reschedule or cancel appointment
func (s *Server) UpdateAppointment(context echo.Context) error {
	log.Println("Update appointment")
	body, err := parseutil.ParseJSON(context)
	if err != nil {
		log.Printf("\nError: %+v", err)
	}
	//
	fmt.Println(body)
	required := []string{} //slotDateTime, cancelled
	remove := []string{"accountId", "clinicId", "patientId", "contactPhone"}
	body = parseutil.RemoveFields(body, remove)
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	body["accountId"] = accountID
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return context.JSON(http.StatusBadRequest, missing)
	}

	stringFields := []string{"slotDateTime"}
	boolField := []string{"cancelled"}
	intField := []string{}
	floatField := []string{}
	jsonField := []string{}
	appointment := model.Appointment{}

	body, invalidType := parseutil.MapX(body, appointment, stringFields, floatField, intField, boolField, jsonField)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return context.JSON(http.StatusBadRequest, invalidType)
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	// query, values := querybuilder.BuildInsertQuery(body, "appointment")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	fmt.Println(body)
	query, values := querybuilder.BuildUpdateQuery(body, "appointment")
	// validation
	if body["slotDateTime"] != nil {
		//layout := "2006-01-02T15:04:05.000Z"
		strDate := body["slotDateTime"].(string)
		t, err := time.Parse(time.RFC3339, strDate)

		if err != nil {
			return context.JSON(http.StatusInternalServerError, err)
		}
		if t.Sub(time.Now()) < 0 {
			return context.JSON(http.StatusInternalServerError, errors.New("Future time is needed"))
		}

	}

	query = query + fmt.Sprintf(" WHERE uid = $%d AND account_id = $%d", len(values)+1, len(values)+2)
	query = query + " RETURNING uid, account_id, clinic_id, patient_id, disease_id, slot_date_time, contact_phone, no_show, created_on, cancelled"
	fmt.Println(query)

	values = append(values, context.Param("uid"))
	values = append(values, accountID)
	// Execute query
	tx := s.DB.DBORM.Begin()
	if tx.Error != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, tx.Error)
	}
	row := tx.Raw(query, values)
	//diseaseID := sql.NullInt64{}
	fmt.Println(query)
	fmt.Println(values)

	row.Scan(&appointment)
	if tx.Error != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, tx.Error)
	}
	tx.Commit()
	if tx.Error != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return context.JSON(http.StatusInternalServerError, tx.Error)
	}

	// Parse response into {model.User}: ParseRow(row, returnfields)
	return context.JSON(http.StatusOK, appointment)

}

//UpdateTreatment update diagnosis, prescription & test
func UpdateTreatment(context echo.Context) error {
	return nil
}

//ApproveDoctorQualification approve doctor qualification
func ApproveDoctorQualification(context echo.Context) error {
	return nil
}

//ApproveDoctorProfile approve doctore profile after document verification
func ApproveDoctorProfile(context echo.Context) error {
	return nil
}

// ApproveClinic approve clinic
func ApproveClinic(context echo.Context) error {
	return nil
}

// ActivateDoctorProfile activate or deactivate doctor
func ActivateDoctorProfile(context echo.Context) error {
	return nil
}

//UpdateDoctorReview  should we allow user to allow review?
