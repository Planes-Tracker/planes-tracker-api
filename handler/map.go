package handler

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/LockBlock-dev/planes-tracker-api/constant"
	apiContext "github.com/LockBlock-dev/planes-tracker-api/context"
	"github.com/LockBlock-dev/planes-tracker-api/database/scope"
	"github.com/LockBlock-dev/planes-tracker-api/dto"
	"github.com/LockBlock-dev/planes-tracker-api/model"
	"github.com/LockBlock-dev/planes-tracker-api/utils"
	"github.com/labstack/echo/v4"
)

type HeatmapFlightPoint struct {
	Latitude  *float32 `gorm:"column:latitude" json:"latitude"`
	Longitude *float32 `gorm:"column:longitude" json:"longitude"`
}

type TrailsFlightPoint struct {
	FlightId  uint     `json:"flightId"`
	Latitude  *float32 `gorm:"column:latitude" json:"latitude"`
	Longitude *float32 `gorm:"column:longitude" json:"longitude"`
	Altitude  *int32   `gorm:"column:altitude" json:"altitude"`
}

func (h *Handler) GetMapSettings(c echo.Context) error {
	lat, _ := strconv.ParseFloat(os.Getenv("TRACKER_LOCATION_LATITUDE"), 64)
	lon, _ := strconv.ParseFloat(os.Getenv("TRACKER_LOCATION_LONGITUDE"), 64)
	radius, _ := strconv.Atoi(os.Getenv("TRACKER_RADIUS_DISTANCE"))

	box := utils.NewBoundingBox(lat, lon, float64(radius))

	return c.JSON(
		http.StatusOK,
		dto.NewResponse(dto.NewMapSettings(lat, lon, &utils.Radius{Distance: uint32(radius)}, box)),
	)
}

func (h *Handler) GetHeatmapPoints(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.RESPONSE_DEFAULT_TIMEOUT)
	query := h.DB.WithContext(ctx)
	var count int64 = 0
	var flightPoints []HeatmapFlightPoint

	defer cancel()

	query.Model(&model.FlightPoint{}).Count(&count)

	query = query.Model(&model.FlightPoint{}).Find(&flightPoints)

	if query.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.NewError(query.Error.Error()),
		)
	}

	var points dto.HeatmapCoordinateCollection
	for _, point := range flightPoints {
		if point.Latitude != nil && point.Longitude != nil {
			points = append(
				points,
				dto.HeatmapCoordinatesTuple{
					*point.Latitude,
					*point.Longitude,
				},
			)
		}
	}

	return c.JSON(
		http.StatusOK,
		dto.NewCollectionReponse(points, count),
	)
}

func (h *Handler) GetTrails(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.RESPONSE_DEFAULT_TIMEOUT)
	query := h.DB.WithContext(ctx)
	var count int64 = 0
	var flightPoints []TrailsFlightPoint

	defer cancel()

	pc := c.(*apiContext.PaginationContext)

	query.Model(&model.FlightPoint{}).Count(&count)

	query = query.Model(
		&model.FlightPoint{},
	).Scopes(
		scope.Paginate(pc.Page(), pc.PageSize()),
	).Scopes(
		scope.Order(constant.ORDERING_DESCENDING, "created_at"),
	).Find(
		&flightPoints,
	)

	if query.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.NewError(query.Error.Error()),
		)
	}

	groupedFlightPoints := make(dto.TrailsFlightPointCollection)
	for _, point := range flightPoints {
		if point.Latitude != nil && point.Longitude != nil {
			var altitude float32 = -1.0

			if point.Altitude != nil {
				altitude = float32(*point.Altitude)
			}

			key := strconv.Itoa(int(point.FlightId))

			groupedFlightPoints[key] = append(
				groupedFlightPoints[key],
				dto.TrailsFlightPoint{
					*point.Latitude,
					*point.Longitude,
					altitude,
				},
			)
		}
	}

	return c.JSON(
		http.StatusOK,
		dto.NewCollectionReponse(groupedFlightPoints, count),
	)
}
