package model

import "time"

type Flight struct {
	FlightId  uint      `gorm:"primarykey" json:"flightId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Registration *string `gorm:"column:registration;uniqueIndex:idx_unique_flight" json:"registration"`
	Flight       *string `gorm:"column:flight" json:"flight"`
	Callsign     *string `gorm:"column:callsign;uniqueIndex:idx_unique_flight" json:"callsign"`
	Origin       *string `gorm:"column:origin" json:"origin"`
	Destination  *string `gorm:"column:destination" json:"destination"`
	DivertedTo   *string `gorm:"column:diverted_to" json:"divertedTo"`
	Model        *string `gorm:"column:model" json:"model"`
	ICAOAddress  *string `gorm:"column:icao_address;type:char(6);uniqueIndex:idx_unique_flight" json:"icaoAddress"`
}
