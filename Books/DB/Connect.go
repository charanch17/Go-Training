package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Establish_connection(config *Config) (*gorm.DB, error) {
	// DATA SOURCE NAME
	var connectionString string
	var db *gorm.DB
	var err error

	switch config.DBMS {
	case "postgres":
		connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.UserName, config.Password, config.Host, config.Port, config.DataBase)
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	}
	return db, err

}
