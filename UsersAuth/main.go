package main

import (
	"log"
	"net/http"
	"os"

	controllers "auth.com/charan/authenticate/Controllers"
	db "auth.com/charan/authenticate/DB"
	initializers "auth.com/charan/authenticate/Initializers"
	models "auth.com/charan/authenticate/Models"
	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
	config := db.Configurtaion{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		DBName:   os.Getenv("DB_NAME"),
		DBMS:     os.Getenv("DBMS_NAME"),
		UserName: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	Database, err := initializers.EstablishDbConnection(&config)
	if err != nil {
		panic(err)
	} else {
		tables := []interface{}{&models.User{}}
		initializers.SyncDbTables(Database, tables)
		db.Instance = Database
	}

}

func main() {
	router := mux.NewRouter()
	controller := controllers.New(db.Instance)
	router.HandleFunc("/signup", controller.Signup).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	router.HandleFunc("/validateToken", controller.Validate).Methods("GET")
	router.HandleFunc("/refreshToken", controller.RefreshToken).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
