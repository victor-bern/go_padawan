package seed

import (
	"go_padawan/src/models"

	"gorm.io/gorm"
)

var currencies []models.Currency

func SeedCurrency(db *gorm.DB) {
	var crc models.Currency
	db.Find(&crc, "1")
	if len(crc.Type) > 1 {
		return
	}
	currencies = []models.Currency{
		{
			Type:   "BRL",
			Symbol: "R$",
		},
		{
			Type:   "USD",
			Symbol: "$",
		},
		{
			Type:   "EUR",
			Symbol: "€",
		},
		{
			Type:   "BTC",
			Symbol: "B",
		},
	}
	for _, v := range currencies {
		db.Create(&v)
	}
}
