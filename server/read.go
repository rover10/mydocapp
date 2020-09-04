package server

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rover10/mydocapp.git/model"
	"github.com/rover10/mydocapp.git/response"
	"github.com/rover10/mydocapp.git/token"
)

//Doctor read a doctor
func Doctor(context echo.Context) error {
	return nil
}

//Clinic read a clinic detail
func (s *Server) Clinic(context echo.Context) error {
	log.Println("Clinic")
	clinics := []model.Clinic{}
	if err := s.DB.DBORM.Table("clinic").Find(&clinics).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, clinics)
}

//Clinic read a clinic detail
func (s *Server) GenerateQr(context echo.Context) error {
	//login := token.GetLoggedIn(context)
	//accountID := login["uid"].(string)
	log.Println("GenerateQr")
	if appointmentID := context.QueryParam("appointmentId"); appointmentID != "" {
		fmt.Println(appointmentID)
		qrCode, _ := GenerateQrCode(appointmentID, 128, 128)
		return png.Encode(context.Response().Writer, qrCode)
	}
	return context.JSON(http.StatusNoContent, "")
	//return context.JSON(http.StatusOK, clinics)
}

func GenerateQrCode(data string, l int, w int) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(data, qr.L, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, l, w)
}

//Appointment read appointment
func (s *Server) Appointment(context echo.Context) error {
	log.Println("Appointment")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	fmt.Println(accountID)
	appointments := []model.Appointment{}
	if err := s.DB.DBORM.Table("appointment").Where("account_id = ?", accountID).Find(&appointments).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, appointments)
	//s.DB.RetrieveAppointment(sid)

}

//Appointment read appointment
func (s *Server) AppointmentV2(context echo.Context) error {
	log.Println("AppointmentV2")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	fmt.Println(accountID)
	appointments := []response.Appointment{}

	cancelledS := context.QueryParam("isCancelled")
	historyS := context.QueryParam("history")
	isCancelled := false
	history := false

	if cancelledS != "" {
		if strings.ToLower(cancelledS) == "true" {
			isCancelled = true
		}
	}

	if historyS != "" {
		if strings.ToLower(historyS) == "true" {
			history = true
		}
	}
	//.Preload("Clinic").Preload("Patient").Preload("Doctor")
	//clinic := response.Clinic{}
	ret := s.DB.DBORM.Debug().
		//Preload("Clinic").
		Table("appointment").
		Preload("Clinic", func(db *gorm.DB) *gorm.DB {
			return db.
				Table("clinic")
		}).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Table("patient")
		}).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Table("doctor")
		})

	//	Where("account_id = ?", accountID)
	if isCancelled {
		ret = ret.Where("account_id = ? AND cancelled = 'True'", accountID)
	} else if history {
		ret = ret.Where("account_id = ? AND slot_date_time < now()", accountID)
	} else {
		ret = ret.Where("account_id = ? AND slot_date_time > now()", accountID)
	}

	if err := ret.Find(&appointments).Error; err != nil {
		return err
	}

	//s.DB.DBORM.
	return context.JSON(http.StatusOK, appointments)
	//s.DB.RetrieveAppointment(sid)

}

//Patient read linked patient
func (s *Server) Patient(context echo.Context) error {

	log.Println("Patient")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	patients := []model.Patient{}
	if err := s.DB.DBORM.Table("patient").Where("account_id = ?", accountID).Find(&patients).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, patients)
}

//Treatment read treatment detail which includes patient_problem_description, prescription, test
func Treatment(context echo.Context) error {
	return nil
}
