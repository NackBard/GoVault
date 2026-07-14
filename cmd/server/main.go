package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/NackBard/GoVault/internal/handler"
	"github.com/NackBard/GoVault/internal/store"
	_ "github.com/mattn/go-sqlite3"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	slog.Info("health")
	fmt.Fprintf(w, `{"status": "ok"}`)
}

func main() {
	db, err := sql.Open("sqlite3", "./govault.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrationQuery, err := os.ReadFile("migrations/001_init.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(migrationQuery))
	if err != nil {
		log.Fatal(err)
	}

	noteStore := store.NewSQLiteNoteStore(db)
	noteHandler := handler.NewNoteHandler(noteStore)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /notes", noteHandler.HandleList)
	mux.HandleFunc("POST /notes", noteHandler.HandleCreate)
	mux.HandleFunc("GET /notes/{id}", noteHandler.HandleGetByID)
	mux.HandleFunc("PUT /notes/{id}", noteHandler.HandleUpdate)
	mux.HandleFunc("DELETE /notes/{id}", noteHandler.HandleDelete)
	mux.HandleFunc("GET /health", healthHandler)

	slog.Info("server started", "port", 8080)
	if err := http.ListenAndServe(":8080", handler.LoggingMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
