package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping - This function will ping the echo server
func Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Health": "OK"})
}
