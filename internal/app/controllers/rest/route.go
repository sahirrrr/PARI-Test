package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	ServicesImpl struct {
		Services
	}

	Services struct {
		dig.In
	}
)

func SetRoute(
	e *echo.Echo,
	services Services,
) {
	// handler := ServicesImpl{Services: services}

	// Health
	e.GET("/health", healthCheck)

	// Items

}

func healthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, "OK")
}
