package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) Error(c echo.Context, err error) error {
	return c.Render(http.StatusInternalServerError, "failed", ErrorData{
		ErrorMessage: "Unknown Error Occurred",
		ErrorDetails: err.Error(),
	})
}
