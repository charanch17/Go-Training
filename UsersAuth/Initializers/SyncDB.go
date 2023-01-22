package initializers

import "gorm.io/gorm"

func SyncDbTables(db *gorm.DB, tables []interface{}) error {
	return db.AutoMigrate(tables...)
}
