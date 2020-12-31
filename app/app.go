package app

import (
	"log"

	"github.com/harmony-development/inviter/app/config"
	"github.com/harmony-development/inviter/app/manager"
	"github.com/harmony-development/inviter/app/templateengine"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	Manager *manager.Manager
	Config  *config.Config
}

// Start starts the app
func Start(m *manager.Manager) {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := server{m, c}

	e := echo.New()
	e.Renderer = templateengine.Template{}

	e.Use(middleware.Recover())

	e.Static("/static", "./app/static")

	e.GET("/_media/:homeserver/:attachment_id", s.attachmentProxy)
	e.GET("/:homeserver/:invite_id", s.foreignInvite)
	e.GET("/:invite_id", s.defaultInvite)
	e.GET("/", s.index)

	e.Start(":2290")
}
