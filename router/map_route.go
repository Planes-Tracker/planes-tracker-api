package router

import (
	"github.com/LockBlock-dev/planes-tracker-api/handler"
	"github.com/labstack/echo/v4"
)

func MapRoute(e *echo.Echo, h *handler.Handler) {
	e.GET("/map/settings", h.GetMapSettings)
	e.GET("/map/heatmap", h.GetHeatmapPoints)
	e.GET("/map/trails", h.GetTrails)
}
