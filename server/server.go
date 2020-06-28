package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/rover10/mydocapp.git/config"
	"github.com/rover10/mydocapp.git/database"
)

type Server struct {
	config.Config
	Router          *echo.Echo
	Templates       *template.Template
	DB              *database.DocDB
	SECRET_PASSWORD string
}

var SECRET_PASSWORD = "Ra@@ndom&%#@%(%5*&%^&Value(&*HJGJGJggHHJKJBJ"

// Start - This function will start the echo server
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	log.Infof("Listening on %s", address)
	return s.Router.Start(address)
}

// Stop - This function will stop the echo server
func (s *Server) Stop(ctx context.Context) error {
	return s.Router.Shutdown(ctx)
}

// ServeHTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SERVING..... ...")
	s.Router.ServeHTTP(w, r)
}

func loadTemplates(webDir string) *template.Template {
	templatePath := path.Join(webDir, "templates", "*.html")
	return template.Must(template.New("").Delims("[[", "]]").ParseGlob(templatePath))
}

// Render - Render the HTML to echo server
func (s *Server) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return s.Templates.ExecuteTemplate(w, name, data)
}

func Templates(webDir string) *template.Template {
	templatePath := path.Join(webDir, "templates", "*.html")
	return template.Must(template.New("").Delims("[[", "]]").ParseGlob(templatePath))
}

// NewServer - Constructor function for server
func NewServer(cfg config.Config) *Server {
	server := &Server{
		Config:          cfg,
		Router:          echo.New(),
		Templates:       Templates("web"),
		DB:              nil,
		SECRET_PASSWORD: "Ra@@ndom&%#@%(%5*&%^&Value(&*HJGJGJggHHJKJBJ",
	}
	return server
}
