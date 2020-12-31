package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) foreignInvite(c echo.Context) error {
	data, err := s.Manager.Get(c.Param("homeserver"), c.Param("invite_id"))
	if err != nil {
		return nil
	}
	return c.Render(http.StatusOK, "invite", data)
}
