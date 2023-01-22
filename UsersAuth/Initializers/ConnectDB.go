package initializers

import (
	"fmt"

	db "auth.com/charan/authenticate/DB"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EstablishDbConnection(config *db.Configurtaion) (*gorm.DB, error) {
	ConnectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.DBMS, config.UserName, config.Password, config.Host, config.Port, config.DBName)
	return gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})
}
