package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greeter() {
	fmt.Println("hello world")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>helo</h1>"))

}

func main() {

	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
