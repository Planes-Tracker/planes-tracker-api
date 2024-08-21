package dto

import "github.com/LockBlock-dev/planes-tracker-api/utils"

type MapSettings struct {
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	Radius    int         `json:"radius"`
	Bounds    [][]float64 `json:"bounds"`
}

func NewMapSettings(lat float64, lon float64, radius *utils.Radius, box *utils.BoundingBox) *MapSettings {
	return &MapSettings{
		Latitude:  lat,
		Longitude: lon,
		Radius:    int(radius.AsMeters()),
		Bounds: [][]float64{
			{box.NorthWest.Lat, box.NorthWest.Lon},
			{box.SouthEast.Lat, box.SouthEast.Lon},
		},
	}
}

type HeatmapCoordinatesTuple = [2]float32

type HeatmapCoordinateCollection = []HeatmapCoordinatesTuple

type TrailsFlightPoint = [3]float32

type TrailsFlightPointCollection = map[string][]TrailsFlightPoint
