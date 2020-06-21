package api

import (
	"path"

	"github.com/labstack/echo/middleware"
	"github.com/rover10/mydocapp.git/server"
)

func Api(server *server.Server) {
	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())
	server.Router.Use(middleware.CORS())
	server.Router.GET(path.Join(server.APIPath, "/v1/ping"), server.Ping)
	server.Router.GET(path.Join(server.APIPath, "/v1/user/:uid"), server.RegisterUser)
	server.Router.POST(path.Join(server.APIPath, "/v1/user"), server.RegisterUser)
	server.Router.POST(path.Join(server.APIPath, "/v1/doctor"), server.RegisterDoctor)

	server.Router.Renderer = server
	server.Router.HideBanner = true
	server.Router.HidePort = true
}
