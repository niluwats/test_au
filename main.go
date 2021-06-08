package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("started")
	router := mux.NewRouter()
	router.HandleFunc("/auth/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode("app started"); err != nil {
			panic(err)
		}
	}).Methods(http.MethodGet)
	address := "localhost"
	port := "8000"
	server := fmt.Sprintf("%s:%s", address, port)
	log.Fatal(http.ListenAndServe(server, router))
}
