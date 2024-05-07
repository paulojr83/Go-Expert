package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

var Dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func GetDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(Dictionary)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PostDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var dict Dict

	err := json.NewDecoder(r.Body).Decode(&dict)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := Dictionary[dict.Key]; ok {
		w.WriteHeader(http.StatusConflict)
	} else {
		Dictionary[dict.Key] = dict.Value
	}
	err = json.NewEncoder(w).Encode(Dictionary)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Rather calling http.HandleFunc, call the equivalent router.HandleFunc
	// This gives us access to method-based routing
	router.HandleFunc("/dictionary", GetDictionary).Methods(http.MethodGet)
	router.HandleFunc("/dictionary", PostDictionary).Methods(http.MethodPost)
	fmt.Println("Server is starting on port 3000...")
	// The second argument for "ListenAndServe" was previously nil
	/// Now that we are using our own custom router, we pass it along to "ListenAndServe" as its second argument
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
