package api

import (
	"path"

	"github.com/labstack/echo/middleware"
	"github.com/rover10/mydocapp.git/handler"
	"github.com/rover10/mydocapp.git/server"
)

func Api(server *server.Server) {
	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())
	server.Router.Use(middleware.CORS())
	server.Router.GET(path.Join(server.APIPath, "/api/v1/ping"), handler.Ping)
	server.Router.Renderer = server
	server.Router.HideBanner = true
	server.Router.HidePort = true
}
