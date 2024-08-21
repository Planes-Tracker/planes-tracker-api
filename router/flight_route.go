package router

import (
	"github.com/LockBlock-dev/planes-tracker-api/handler"
	"github.com/labstack/echo/v4"
)

func FlightRoute(e *echo.Echo, h *handler.Handler) {
	e.GET("/flight/:flightId", h.GetFlight)
	e.GET("/flights", h.GetAllFlights)
	e.GET("/flights/count", h.GetFlightsCount)
}
