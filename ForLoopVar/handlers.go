package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world :]</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	// This handler writes HTML to w
	fmt.Fprintf(w, "<h1>Contact Info</h1>")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	cities := []string{"Acrelândia",
		"Assis Brasil",
		"Brasiléia",
		"Bujari",
		"Capixaba",
		"Cruzeiro do Sul",
		"Epitaciolândia",
		"Feijó",
		"Jordão",
		"Mâncio Lima",
		"Manoel Urbano",
		"Marechal Thaumaturgo",
		"Plácido de Castro",
		"Porto Acre",
		"Porto Walter",
		"Rio Branco",
		"Rodrigues Alves",
		"Santa Rosa do Purus",
		"Sena Madureira",
		"Senador Guiomard",
		"Tarauacá",
		"Xapuri"}

	var list string
	for _, city := range cities {
		list += fmt.Sprintf("City name: %s\n", city)
	}

	fmt.Fprintf(w, list)
}

var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(dictionary)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
func main() {

	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /contact", contact)

	http.HandleFunc("GET /cities", cityList)

	http.HandleFunc("GET /dictionary", getDictionary)

	fmt.Println("Server running...")
	http.ListenAndServe(":3000", nil)
}
