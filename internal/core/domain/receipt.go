package domain

import "time"

type Receipt struct {
	Id int `json:"id"`
	Version int `json:"version"`

	UserId int `json:"user_id"`
	Name string `json:"name" example:"Чек из вкусвилла"`
	Amount float64 `json:"amount" example:"2499.99"`
	CreatedAt time.Time `json:"created_at" example:"2026-09-12T10:42:52Z"`
}

type ReceiptRequest struct {
	Name string `json:"name" example:"Чек из вкусвилла"`
	Amount float64 `json:"amount" example:"2499.99"`
} // @name ReceiptRequest