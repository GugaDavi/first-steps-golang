package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Books collection
var Books = []Book{
	{
		ID:    0,
		Title: "Harry Potter",
		Owner: "JK Roling",
	},
	{
		ID:    1,
		Title: "Harry Potter 2",
		Owner: "JK Roling",
	},
}

func main() {
	serverConfiguration()
}

func serverConfiguration() {
	port := 3333

	routes()

	println("Servin listining on port:", port)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

type routeFunc func(w http.ResponseWriter, r *http.Request)

func routes() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		routesConfig(rw, r, mainRouter)
	})
	http.HandleFunc("/books", func(rw http.ResponseWriter, r *http.Request) {
		routesConfig(rw, r, bookRouter)
	})
	http.HandleFunc("/books/", func(rw http.ResponseWriter, r *http.Request) {
		routesConfig(rw, r, bookRouter)
	})
}

func routesConfig(w http.ResponseWriter, r *http.Request, f routeFunc) {
	w.Header().Set("Content-Type", "application/json")
	f(w, r)
}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func bookRouter(w http.ResponseWriter, r *http.Request) {
	splices := strings.Split(r.URL.Path, "/")

	if len(splices) == 2 || len(splices) == 3 && splices[2] == "" {
		if r.Method == "GET" {
			getBooks(w, r)
		} else if r.Method == "POST" {
			createBook(w, r)
		}
	} else if len(splices) == 3 || len(splices) == 4 && splices[3] == "" {
		if r.Method == "GET" {
			getBook(w, r)
		} else if r.Method == "DELETE" {
			deleteBook(w, r)
		}
	} else {
		fmt.Fprintf(w, "Endpoint not found")
		w.WriteHeader(http.StatusNotFound)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	splited := strings.Split(r.URL.Path, "/")

	if len(splited) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var id, e = strconv.Atoi(splited[2])
	var book Book

	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, b := range Books {
		if b.ID == id {
			book = b
		}
	}

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	splited := strings.Split(r.URL.Path, "/")

	if len(splited) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var id, e = strconv.Atoi(splited[2])

	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bookIndex := -1
	for i, book := range Books {
		if book.ID == id {
			bookIndex = i
		}
	}

	if bookIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Books = append(Books[0:bookIndex], Books[bookIndex+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	body, _ := ioutil.ReadAll(r.Body)

	var newBook Book
	json.Unmarshal(body, &newBook)
	newBook.ID = len(Books)
	Books = append(Books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// Book Model
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
}
