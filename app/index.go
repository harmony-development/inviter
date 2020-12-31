package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
