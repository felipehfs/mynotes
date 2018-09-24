package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/mynotes/config"
)

type DBContext string

// WithDB setup the database to next handlers
func WithDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		database, err := config.GetDBInstance()

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Fatalln(err)
			return
		}
		ctx := context.WithValue(r.Context(), DBContext("database"), database)
		request := r.WithContext(ctx)
		next.ServeHTTP(w, request)
	})
}

// ExtractDB returns the database of request
func ExtractDB(w http.ResponseWriter, r *http.Request) (*config.Database, error) {
	db, ok := r.Context().Value(DBContext("database")).(*config.Database)
	if !ok {
		return nil, errors.New("Context or database wrong")
	}
	return db, nil
}
