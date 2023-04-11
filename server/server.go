package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {

	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")

	fmt.Println("Server is up and running on localhost:8080")
	http.ListenAndServe(":8080", router)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
