package rest

import (
	"hello-go/pkg/core"
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Code    core.ErrorCode `json:"code"`
	Message string         `json:"message"`
}

// HTTPErrorHandler for app
func HTTPErrorHandler(err error, c echo.Context) {
	if err, ok := err.(*core.Error); ok {
		r := errorResponse{Code: err.Code, Message: err.Message}

		switch err.Type {
		case core.InternalServerError:
			c.JSON(http.StatusInternalServerError, r)

		case core.ValidationError:
			c.JSON(http.StatusBadRequest, r)
		}
	}
}

// RegisterErrorHandler registers error handler to echo
func RegisterErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = HTTPErrorHandler
}
