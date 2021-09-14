package models

import "time"

type Currency struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Type      string    `json:"from"`
	Symbol    string    `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateddAt"`
}
