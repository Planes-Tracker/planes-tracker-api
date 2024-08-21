package main

import (
	"github.com/LockBlock-dev/planes-tracker-api/database"
	"github.com/LockBlock-dev/planes-tracker-api/handler"
	"github.com/LockBlock-dev/planes-tracker-api/middleware"
	"github.com/LockBlock-dev/planes-tracker-api/router"
	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	verboseDB := false
	db := database.ConnectDB(verboseDB)

	app.Use(middleware.Pagination())
	app.Use(eMiddleware.Logger())

	handler := &handler.Handler{DB: db}

	router.FlightRoute(app, handler)
	router.MapRoute(app, handler)

	app.Logger.Fatal(app.Start(":8000"))
}
