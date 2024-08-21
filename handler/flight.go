package handler

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/LockBlock-dev/planes-tracker-api/constant"
	apiContext "github.com/LockBlock-dev/planes-tracker-api/context"
	"github.com/LockBlock-dev/planes-tracker-api/database/scope"
	"github.com/LockBlock-dev/planes-tracker-api/database/scope/filter"
	"github.com/LockBlock-dev/planes-tracker-api/dto"
	"github.com/LockBlock-dev/planes-tracker-api/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/schema"
)

func getFlightFields() map[string]string {
	fields := make(map[string]string)
	v := reflect.ValueOf(model.Flight{})
	s, _ := schema.Parse(&model.Flight{}, &sync.Map{}, schema.NamingStrategy{})

	for i := 0; i < v.Type().NumField(); i++ {
		fields[v.Type().Field(i).Tag.Get("json")] = s.DBNames[i]
	}

	return fields
}

var jsonToDBName = getFlightFields()

func (h *Handler) GetFlight(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.RESPONSE_DEFAULT_TIMEOUT)
	flightId := c.Param("flightId")
	var flight model.Flight

	defer cancel()

	result := h.DB.WithContext(ctx).First(&flight, flightId)

	if result.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.NewError(result.Error.Error()),
		)
	}

	return c.JSON(
		http.StatusOK,
		dto.NewResponse(flight),
	)
}

func (h *Handler) GetAllFlights(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.RESPONSE_DEFAULT_TIMEOUT)
	query := h.DB.WithContext(ctx)
	var count int64 = 0
	var flights []model.Flight

	defer cancel()

	pc := c.(*apiContext.PaginationContext)

	switch strings.ToLower(pc.Filter()) {
	case strings.ToLower(constant.FILTERS_MILITARY):
		query = query.Scopes(filter.Military)
	}

	query.Model(&model.Flight{}).Count(&count)

	sortBy := pc.Sort()
	if sortBy != "" {
		// meh
		if dbName, ok := jsonToDBName[sortBy]; ok {
			sortBy = dbName
		}
	}

	query = query.Scopes(
		scope.Order(pc.Order(), sortBy),
	).Scopes(
		scope.Paginate(pc.Page(), pc.PageSize()),
	).Find(
		&flights,
	)

	if query.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.NewError(query.Error.Error()),
		)
	}

	return c.JSON(
		http.StatusOK,
		dto.NewCollectionReponse(flights, count),
	)
}

func (h *Handler) GetFlightsCount(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.RESPONSE_DEFAULT_TIMEOUT)
	var count int64

	defer cancel()

	result := h.DB.WithContext(ctx).Model(&model.Flight{}).Count(&count)

	if result.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.NewError(result.Error.Error()),
		)
	}

	return c.JSON(
		http.StatusOK,
		dto.NewResponse(count),
	)
}
