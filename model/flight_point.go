package model

import "time"

type FlightPoint struct {
	FlightId      uint `json:"flightId"`
	FlightPointId uint `gorm:"primarykey" json:"flightPointId"`

	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Latitude      *float32  `gorm:"column:latitude" json:"latitude"`
	Longitude     *float32  `gorm:"column:longitude" json:"longitude"`
	Altitude      *int32    `gorm:"column:altitude" json:"altitude"`
	Track         *int32    `gorm:"column:track" json:"track"`
	Speed         *int32    `gorm:"column:speed;default:0" json:"speed"`
	VerticalSpeed *int32    `gorm:"column:vertical_speed" json:"verticalSpeed"`
	OnGround      *bool     `gorm:"column:on_ground;default:false" json:"onGround"`
	SquawkCode    *string   `gorm:"column:squawk" json:"squawk"`
}
