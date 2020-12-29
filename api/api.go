package api

import (
	"path"

	"github.com/labstack/echo/middleware"
	"github.com/rover10/auth"
	"github.com/rover10/server"
)

func Api(server *server.Server) {
	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())
	server.Router.Use(middleware.CORS())
	// https://ednsquare.com/story/jwt-authentication-in-golang-with-echo------T2hTPm
	h := &auth.Handler{}
	server.Router.GET(path.Join(server.APIPath, "/islogin"), server.Ping, auth.IsLoggedIn)
	server.Router.POST("/login", server.Login)
	server.Router.GET("/is-loggedin", h.Private, auth.IsLoggedIn)
	server.Router.GET("/is-admin", h.Private, auth.IsLoggedIn, auth.IsAdmin)
	server.Router.POST("/refresh", h.Token)

	server.Router.GET(path.Join(server.APIPath, "/"), server.Ping)
	server.Router.GET(path.Join(server.APIPath, "/v1/ping"), server.Ping)
	server.Router.GET(path.Join(server.APIPath, "/v1/user/:uid"), server.RegisterUser)
	server.Router.POST(path.Join(server.APIPath, "/v1/user"), server.RegisterUser)
	server.Router.POST(path.Join(server.APIPath, "/v1/doctor"), server.RegisterDoctor, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/clinic"), server.RegisterClinic, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/staff"), server.RegisterStaff, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/patient"), server.RegisterPatient, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/appointment"), server.BookAppointment, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/treatment"), server.RegisterTreatment, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/doctorreview"), server.RegisterDoctorReview, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/staffrole"), server.AssignStaffRole, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/uploaddocument"), server.UploadUserDocument, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/testreport"), server.AddTestReport, auth.IsLoggedIn)
	server.Router.POST(path.Join(server.APIPath, "/v1/uploadqualification"), server.AddDoctorQualification)

	server.Router.GET(path.Join(server.APIPath, "/v1/appointment"), server.Appointment, auth.IsLoggedIn)
	server.Router.GET(path.Join(server.APIPath, "/v2/appointment"), server.AppointmentV2, auth.IsLoggedIn)
	server.Router.PUT(path.Join(server.APIPath, "/v1/appointment/:uid"), server.UpdateAppointment, auth.IsLoggedIn)

	server.Router.GET(path.Join(server.APIPath, "/v1/patients"), server.Patient, auth.IsLoggedIn)
	server.Router.GET(path.Join(server.APIPath, "/v1/clinics"), server.Clinic, auth.IsLoggedIn)
	server.Router.GET(path.Join(server.APIPath, "/v1/qrcode"), server.GenerateQr, auth.IsLoggedIn)
	server.Router.GET(path.Join(server.APIPath, "/v1/search"), server.Search, auth.IsLoggedIn)
	server.Router.File("/", "app/index.html")
	server.Router.Static("static/*", "web/assets")
	//e.Static("/", "assets")
	server.Router.GET(path.Join(server.APIPath, "home"), server.Home)
	//server.Router.File("/home", "public/index.html")
	server.Router.Renderer = server
	server.Router.HideBanner = true
	server.Router.HidePort = true
}
