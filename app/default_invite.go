package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) defaultInvite(c echo.Context) error {
	data, err := s.Manager.Get(s.Config.Server.DefaultServer, c.Param("invite_id"))
	if err != nil {
		return nil
	}
	return c.Render(http.StatusOK, "invite", data)
}
