package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping - This function will ping the echo server
func Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Health": "OK"})
}

func RegiterDoctor(context echo.Context) error {
	return nil
}

func RegisterUser(context echo.Context) error {
	return nil
}

func RegisterClinic(context echo.Context) error {
	return nil
}

func BookAppointment(context echo.Context) error {
	return nil

}

func RegisterMedicalHistory(context echo.Context) error {
	return nil
}

func RegisterPrescription(context echo.Context) error {
	return nil
}
