// go mod tidy -> Dectects what new libraries added to the code and updates it in mod file

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func firstFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("first function of go program"))
}

func secondFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("second function of go program"))
}

//pwd2000@K


func main() {
	c := chi.NewRouter()
	c.Get("/first", firstFunc)
	c.Get("/second", secondFunc)
	err := http.ListenAndServe(":8080", c)
	if err != nil {
		log.Println(err)
	}
}
