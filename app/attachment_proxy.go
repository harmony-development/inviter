package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *server) attachmentProxy(c echo.Context) error {
	host := c.Param("homeserver")
	attachmentID := c.Param("attachment_id")

	if !strings.Contains(host, ":") {
		host = host + ":2289"
	}

	resp, err := http.Get(fmt.Sprintf("https://%s/_harmony/media/download/%s", host, attachmentID))
	if err != nil {
		return err
	}

	_, err = io.Copy(c.Response().Writer, resp.Body)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
