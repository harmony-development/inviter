package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Error(c echo.Context, err error) error {
	status, _ := status.FromError(err)

	displayMessage := "Unknown Error Occurred"
	respStatus := http.StatusInternalServerError

	switch status.Code() {
	case codes.DeadlineExceeded:
		displayMessage = "Timed Out"
		respStatus = http.StatusRequestTimeout
	case codes.Aborted:
		displayMessage = "Host Aborted Connection"
		respStatus = http.StatusServiceUnavailable
	}

	return c.Render(respStatus, "failed", ErrorData{
		ErrorMessage: displayMessage,
	})
}
