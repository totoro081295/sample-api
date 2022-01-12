package main

import (
	"sample-api/database"
	"sample-api/database/migrations"
	"sample-api/database/seeds/sql"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = database.ConnectDB()
	defer database.Close(db)
	migrations.DropTable(db)
	migrations.Migrate(db)

	db.Exec(sql.InsertAccount)
}
