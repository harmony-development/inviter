package app

import (
	"net/http"

	"github.com/harmony-development/inviter/app/manager"
	"github.com/labstack/echo/v4"
)

type ErrorData struct {
	manager.InviteData
	ErrorMessage string
	ErrorDetails string
}

func (s *server) defaultInvite(c echo.Context) error {
	data, err := s.Manager.Get(s.Config.Server.DefaultServer, c.Param("invite_id"))
	if err != nil {
		return s.Error(c, err)
	}
	return c.Render(http.StatusOK, "invite", data)
}
