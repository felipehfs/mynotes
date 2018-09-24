package config

import mgo "gopkg.in/mgo.v2"

type Database struct {
	Client *mgo.Session
}

var db *Database

func newDatabase() (*Database, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Database{Client: session}, nil
}

func (d Database) GetCollection(db, collection string) *mgo.Collection {
	return d.Client.DB(db).C(collection)
}

func GetDBInstance() (*Database, error) {
	if db == nil {
		session, err := newDatabase()
		if err != nil {
			return nil, err
		}
		db = session
	}

	return db, nil
}
