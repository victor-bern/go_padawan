package migrations

import (
	"go_padawan/src/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Exchange{}, models.Currency{})
}
