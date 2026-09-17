package domain

import "time"

type Expense struct {
	Id int `json:"id"`
	Version int `json:"version"`

	Name string `json:"name"`
	Amount float64 `json:"amount"`
	ReceiptId int `json:"receipt_id"`
	UserId int `json:"user_id"`
	CategoryId int `json:"category_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ExpenseRequest struct {
	Name string `json:"name"`
	Amount int `json:"amount"`
	CategoryId int `json:"category_id"`
} // @name ExpenseRequest