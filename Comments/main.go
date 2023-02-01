package main

import (
	"log"
	"net/http"

	handlers "comments.com/charan/comments/Handlers"
	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()
	Router.HandleFunc("/comments", handlers.GetComments).Methods("GET")
	Router.HandleFunc("/comments/commentor/{commentorid}", handlers.GetComments).Methods("GET")
	Router.HandleFunc("/comments/post/{postid}", handlers.GetComments).Methods("GET")
	Router.HandleFunc("/comments/add", handlers.AddComment).Methods("POST")
	Router.HandleFunc("/comments/update", handlers.UpdateComment).Methods("PUT")
	Router.HandleFunc("/comments/delete", handlers.DeleteComment).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
