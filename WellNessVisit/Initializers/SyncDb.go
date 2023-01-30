package initializers

import db "Privia.com/charan/WellNessVisits/DB"

func SyncDb(tables []interface{}) {
	db.DbInst.AutoMigrate(tables...)
}
