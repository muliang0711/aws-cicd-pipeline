package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handleItems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		items, err := listItems(ctx, db)
		if err != nil {
			log.Printf("Read items from MySQL: %v", err)
			http.Error(w, "Unable to load items", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func listItems(ctx context.Context, db *sql.DB) ([]item, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name FROM items ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]item, 0)
	for rows.Next() {
		var record item
		if err := rows.Scan(&record.ID, &record.Name); err != nil {
			return nil, err
		}
		items = append(items, record)
	}
	return items, rows.Err()
}
