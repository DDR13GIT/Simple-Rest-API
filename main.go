package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Golang", Author: "Google", Price: 200.00})
	books = append(books, Book{ID: "2", Title: "Python", Author: "Guido van Rossum", Price: 300.00})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request Payload", http.StatusBadRequest)
	}
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["id"] {
			var updatedBook Book
			if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
				http.Error(w, "Invalid request Payload", http.StatusBadRequest)
				return
			}

			updatedBook.ID = item.ID
			books[i] = updatedBook

			json.NewEncoder(w).Encode(updatedBook)
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
            json.NewEncoder(w).Encode(item)
            return
		}
	}

    http.Error(w, "Book not found", http.StatusNotFound)
}
