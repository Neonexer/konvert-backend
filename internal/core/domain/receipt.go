package domain

import "time"

type Receipt struct {
	Id int `json:"id"`
	Version int `json:"version"`

	UserId int `json:"user_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Amount float64 `json:"amount"`
}

type ReceiptRequest struct {
	Name string `json:"name"`
	Amount float64 `json:"amount"`
} // @name ReceiptRequest