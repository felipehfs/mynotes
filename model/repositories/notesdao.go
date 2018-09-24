// Package repositories contains all operations like queries,
// insert, delete of data
package repositories

import (
	"github.com/mynotes/config"
	"github.com/mynotes/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NotesDao is in the charge of all database operations
type NotesDao struct {
	Database *config.Database
}

func (nd NotesDao) getnotes() *mgo.Collection {
	return nd.Database.GetCollection("blasternotes", "notes")
}

// Create push new data to database
func (nd NotesDao) Create(note model.Note) error {
	return nd.getnotes().Insert(note)
}

// Read retrieves all notes from database
func (nd NotesDao) Read() ([]model.Note, error) {
	var notes []model.Note
	err := nd.getnotes().Find(bson.M{}).All(&notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// Delete drop the note
func (nd NotesDao) Delete(id string) error {
	return nd.getnotes().Remove(bson.M{"id": bson.ObjectIdHex(id)})
}

// Update change the note indicates by id
func (nd NotesDao) Update(id string, note model.Note) error {
	return nd.getnotes().Update(bson.M{"id": bson.ObjectIdHex(id)}, note)
}

// Find retrieves the note by ID
func (nd NotesDao) Find(id string) (*model.Note, error) {
	var note model.Note

	err := nd.getnotes().Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}
