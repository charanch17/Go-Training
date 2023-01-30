package main

import (
	"log"
	"net/http"
	"os"

	db "Privia.com/charan/WellNessVisits/DB"
	handlers "Privia.com/charan/WellNessVisits/Handlers"
	initializers "Privia.com/charan/WellNessVisits/Initializers"
	middleware "Privia.com/charan/WellNessVisits/MiddleWare"
	models "Privia.com/charan/WellNessVisits/Models"
	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
	config := &db.Config{
		UserName: os.Getenv("DB_USER_NAME"),
		DataBase: os.Getenv("DB_NAME"),
		DBMS:     os.Getenv("DBMS"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}
	instance, err := db.EstablishDbConnection(config)
	if err != nil {
		panic(err)
	}
	db.DbInst = instance
	tables := []interface{}{&models.User{}, &models.Role{}, &models.Visits{}}
	initializers.SyncDb(tables)

}
func main() {
	Router := mux.NewRouter()
	handler := handlers.New(db.DbInst)
	Router.HandleFunc("/signup", handler.Signup).Methods("POST")
	Router.HandleFunc("/role/add", middleware.AuthenticateAndCheckPermission(handler.AddRole, []string{"Admin"}, handler)).Methods("POST")
	Router.HandleFunc("/role/delete/{rolename}", middleware.AuthenticateAndCheckPermission(handler.DeleteRole, []string{"Admin"}, handler)).Methods("PUT")
	Router.HandleFunc("/visits", middleware.AuthenticateAndCheckPermission(handler.GetWellNessVisits, []string{"Patient", "Doctor"}, handler)).Methods("GET")
	Router.HandleFunc("/visit/add", middleware.AuthenticateAndCheckPermission(handler.AddWellNessVisit, []string{"Patient"}, handler)).Methods("POST")
	Router.HandleFunc("/login", handler.Login).Methods("POST")
	// Router.HandleFunc("/visits/add", middleware.AuthenticateAndCheckPermission(handler.AddWellNessVisit, []string{"Patient"})).Methods("POST")
	// fmt.Println(db.DbInst)
	log.Fatal(http.ListenAndServe(":8080", Router))
}
