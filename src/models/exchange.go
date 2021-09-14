package models

import "time"

type Exchange struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
	Rate      float64   `json:"rate"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateddAt"`
}
