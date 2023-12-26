package main

import (
	"context"
	"database/sql"
	"log"
)

func saveToDatabase(ctx context.Context, db *sql.DB, currency Currency) error {
	insertSQL := `INSERT INTO currencies(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.ExecContext(ctx,
		insertSQL, currency.Code, currency.Codein, currency.Name, currency.High, currency.Low, currency.VarBid, currency.PctChange, currency.Bid, currency.Ask, currency.Timestamp, currency.Create_date)
	return err
}

func sqlLite() {
	db, err := sql.Open("sqlite3", "currencies.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createTableSQL := `CREATE TABLE IF NOT EXISTS currencies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT,
    codein TEXT,
    name TEXT,
    high TEXT,
    low TEXT,
    varBid TEXT,
    pctChange TEXT,
    bid TEXT,
    ask TEXT,
    timestamp TEXT,
    create_date TEXT
);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
