package services

import (
	"fmt"
	"go_padawan/src/database"
	"go_padawan/src/models"
	"log"
)

func GetCurrency(currency string) (string, error) {
	db := database.GetDatabase()

	var cr models.Currency

	err := db.Where("`type` = ?", currency).First(&cr).Error

	if err != nil {
		return "", fmt.Errorf("cannot find currency")
	}

	return cr.Symbol, nil
}

func ValidateCurrency(from string, to string) bool {
	db := database.GetDatabase()

	var cr models.Currency
	var cr2 models.Currency

	err := db.Where("`type` = ?", from).First(&cr).Error

	if err != nil {
		log.Printf("Primeiro log: %v", err.Error())
		return false
	}

	err = db.Where("`type` = ?", to).First(&cr2).Error

	return err == nil
}
