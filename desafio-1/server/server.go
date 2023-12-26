package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func getFirstCurrency(result map[string]Currency) (string, Currency, bool) {
	for key, currency := range result {
		// Retorna a primeira chave e valor encontrados
		return key, currency, true
	}
	// Retorna valores vazios se o mapa estiver vazio
	return "", Currency{}, false
}

func fetchExchangeRate(ctx context.Context) (Currency, error) {
	var currency Currency

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return currency, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return currency, err
	}
	defer resp.Body.Close()

	var result map[string]Currency
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return currency, err
	}
	
	_, currency, ok := getFirstCurrency(result)
	if ok == false {
		log.Fatal("O mapa está vazio.")
	}
	fmt.Printf("Resultado: %+v\n", currency)
	return currency, nil

}

func cotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		currency, err := fetchExchangeRate(ctx)
		if err != nil {
			log.Println("Error fetching exchange rate:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		ctxSave, cancelSave := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelSave()

		err = saveToDatabase(ctxSave, db, currency)
		if err != nil {
			log.Println("Error saving to database:", err)
			// Não interrompe a execução; apenas loga o erro
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"bid": currency.Bid})
	}
}

func main() {
	sqlLite()
	db, err := sql.Open("sqlite3", "currencies.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", cotacaoHandler(db))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
