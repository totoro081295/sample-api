package migrations

import (
	"sample-api/database"
	"sample-api/domain"

	"gorm.io/gorm"
)

func Execute() {
	db := database.ConnectDB()
	Migrate(db)
	database.Close(db)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&domain.Account{},
	)
}

func DropTable(db *gorm.DB) {
	db.Migrator().DropTable(
		&domain.Account{},
	)
}
