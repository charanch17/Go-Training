package db

import (
	"fmt"

	"golang.org/x/crypto/openpgp/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EstablishDbConnection(config *Config) (*gorm.DB, error) {
	switch config.DBMS {
	case "postgres", "mssql", "mysql":
		ConnectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.DBMS, config.UserName, config.Password, config.Host, config.Port, config.DataBase)
		return gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})
	}
	return nil, new(errors.InvalidArgumentError)
}
