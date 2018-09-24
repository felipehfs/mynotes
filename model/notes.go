// Package model contains all entities
package model

import (
	"encoding/json"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	FORMAT = "01/02/2006"
)

// Note entity mark a task for sheduler
type Note struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	Headline  string        `json:"headline"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created,omitempty"`
	Reminder  time.Time     `json:"reminder,omitempty" bson:"reminder,omitempty"`
}

func (n *Note) UnmarshalJSON(data []byte) error {
	var rawStrings map[string]string
	err := json.Unmarshal(data, &rawStrings)

	if err != nil {
		return err
	}

	for key, value := range rawStrings {
		switch strings.ToLower(key) {
		case "id":
			if value != "" {
				id := bson.ObjectIdHex(value)
				n.ID = id
			} else {
				n.ID = bson.NewObjectId()
			}
		case "headline":
			n.Headline = value
		case "content":
			n.Content = value
		case "created":
			t, err := time.Parse(FORMAT, value)
			if err != nil {
				return err
			}
			n.CreatedAt = t
		case "reminder":
			t, err := time.Parse(FORMAT, value)
			if err != nil {
				return err
			}
			n.Reminder = t
		}
	}
	return nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	basicLink := struct {
		ID        string `json:"id"`
		Headline  string `json:"headline"`
		Content   string `json:"content"`
		CreatedAt string `json:"created"`
		Reminder  string `json:"reminder"`
	}{
		ID:        n.ID.Hex(),
		Headline:  n.Headline,
		Content:   n.Content,
		CreatedAt: n.CreatedAt.Format(FORMAT),
		Reminder:  n.Reminder.Format(FORMAT),
	}

	return json.Marshal(basicLink)
}
