package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Writer *Writer `json:"writer"`
}

type Writer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "1234567", Title: "Harry Potter", Writer: &Writer{FirstName: "Beknur", LastName: "Bekkaliev"}})
	books = append(books, Book{ID: "2", Isbn: "123456", Title: "Harry Potter 2", Writer: &Writer{FirstName: "Beknur2", LastName: "Bekkaliev2"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
