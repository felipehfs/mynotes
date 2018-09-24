package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/mynotes/middleware"
	"github.com/mynotes/model"
	"github.com/mynotes/model/repositories"
)

// CreateNote inserts the new note
func CreateNote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := middleware.ExtractDB(w, r)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Fatal(err)
			return
		}
		dao := repositories.NotesDao{Database: conn}
		var note model.Note

		note.ID = bson.NewObjectIdWithTime(time.Now())
		note.CreatedAt = time.Now()

		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, "Donald trump", http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		defer r.Body.Close()
		if err := dao.Create(note); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
	})
}

// ReadNote retrieves all notes
func ReadNote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := middleware.ExtractDB(w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		dao := repositories.NotesDao{Database: conn}
		notes, err := dao.Read()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	})
}

// DeleteNote deletes the note
func DeleteNote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		conn, err := middleware.ExtractDB(w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		dao := repositories.NotesDao{Database: conn}
		if err := dao.Delete(vars["id"]); err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			log.Fatal(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

// UpdateNote set the changes into database
func UpdateNote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var note model.Note
		conn, err := middleware.ExtractDB(w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		dao := repositories.NotesDao{Database: conn}

		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		if err := dao.Update(vars["id"], note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Panicln(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	})
}

// FindNote searches the note
func FindNote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		conn, err := middleware.ExtractDB(w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		dao := repositories.NotesDao{Database: conn}
		note, err := dao.Find(param["id"])
		if err != nil {
			http.Error(w, "Not found", 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	})
}
