package server

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/labstack/echo"
	"github.com/rover10/mydocapp.git/model"
	"github.com/rover10/mydocapp.git/token"
)

//Doctor read a doctor
func Doctor(context echo.Context) error {
	return nil
}

//Clinic read a clinic detail
func (s *Server) Clinic(context echo.Context) error {
	clinics := []model.Clinic{}
	if err := s.DB.DBORM.Table("clinic").Find(&clinics).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, clinics)
}

//Clinic read a clinic detail
func (s *Server) GenerateQr(context echo.Context) error {
	dataString := "Rakesh"
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 128, 128)
	return png.Encode(context.Response().Writer, qrCode)

	//return context.JSON(http.StatusOK, clinics)
}

//Appointment read appointment
func (s *Server) Appointment(context echo.Context) error {
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

//Patient read linked patient
func (s *Server) Patient(context echo.Context) error {
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
