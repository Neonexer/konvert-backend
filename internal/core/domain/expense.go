package domain

import "time"

type Expense struct {
	Id int `json:"id"`
	Version int `json:"version"`

	Name string `json:"name" example:"Тофу"`
	Amount float64 `json:"amount" example:"260.90"`
	CategoryId int `json:"category_id" example:"23"`
	ReceiptId int `json:"receipt_id"`
	UserId int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ExpenseRequest struct {
	Name string `json:"name" example:"Тофу"`
	Amount float64 `json:"amount" example:"260.90"`
	CategoryId int `json:"category_id" example:"23"`
} // @name ExpenseRequest