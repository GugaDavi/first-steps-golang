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
		routesConfig(rw, r, getBook)
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
	if r.Method == "GET" {
		getBooks(w, r)
	} else if r.Method == "POST" {
		createBook(w, r)
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

	var id, _ = strconv.Atoi(splited[2])
	var book Book

	for i := 0; i < len(Books); i++ {
		if Books[i].ID == id {
			book = Books[i]
		}
	}

	json.NewEncoder(w).Encode(book)
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
