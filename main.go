package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("The program is starting")

	r := mux.NewRouter()
	r.HandleFunc("/monkeys", monkeyHandler)
	r.HandleFunc("/", homeHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func monkeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've hit the monkey routes")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
