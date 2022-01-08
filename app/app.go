package app

import (
	"embed"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/harmony-development/inviter/app/config"
	"github.com/harmony-development/inviter/app/manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	failed *template.Template
	index  *template.Template
	invite *template.Template
}

func New() *Template {
	return &Template{
		failed: template.Must(template.ParseFS(templates, "templates/failed.tmpl", "templates/application.tmpl")),
		index:  template.Must(template.ParseFS(templates, "templates/index.tmpl", "templates/application.tmpl")),
		invite: template.Must(template.ParseFS(templates, "templates/invite.tmpl", "templates/application.tmpl")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	switch name {
	case "failed":
		return t.failed.ExecuteTemplate(w, name+".tmpl", data)
	case "index":
		return t.index.ExecuteTemplate(w, name+".tmpl", data)
	case "invite":
		return t.invite.ExecuteTemplate(w, name+".tmpl", data)
	default:
		return errors.New("template not found")
	}
}

type server struct {
	Manager *manager.Manager
	Config  *config.Config
}

//go:embed "templates"
var templates embed.FS

//go:embed "static"
var static embed.FS

var staticHTTP http.FileSystem

func init() {
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	staticHTTP = http.FS(fsys)
}

// Start starts the app
func Start(m *manager.Manager) {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := server{m, c}

	e := echo.New()
	e.Renderer = New()
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		println(e.Error())
	}

	e.Use(middleware.Recover())

	staticHandler := http.FileServer(staticHTTP)

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", staticHandler)))
	e.GET("/_media/:homeserver/:attachment_id", s.attachmentProxy)
	e.GET("/:homeserver/:invite_id", s.foreignInvite)
	e.GET("/:invite_id", s.defaultInvite)
	e.GET("/", s.index)

	e.Start(":2290")
}
