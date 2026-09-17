package domain

import "time"

type Category struct {
	Id int `json:"id"`
	Version int `json:"version"`

	Name string `json:"name"`
	MaxMonthAmount int `json:"max_month_amount"`
	Color string `json:"color"`
	Icon string `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
}