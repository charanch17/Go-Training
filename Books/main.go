package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "Books.com/charan/books/DB"
	handlers "Books.com/charan/books/Handlers"
	models "Books.com/charan/books/Models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	config := db.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DataBase: os.Getenv("DATABASE"),
		DBMS:     os.Getenv("DBMS"),
	}
	fmt.Println(config.UserName)
	db, err := db.Establish_connection(&config)
	if err != nil {
		panic(err)
	}
	handler := handlers.New(db)
	router := mux.NewRouter()
	db.AutoMigrate(&models.Book{})
	router.HandleFunc("/books", handler.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/add", handler.AddNewBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods(http.MethodDelete)
	router.HandleFunc("/books/update/{id}", handler.UpdataBook).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", router))
}
