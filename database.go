package main

import (
	"encoding/json"
	"os"
	"path"
)

// Database maintains a list of subdomains and paths to use for redirects.
type Database struct {
	name  string
	Paths map[string]string `json:"paths"`
}

// LoadDatabase attempts to load the database from disk. There is a rather
// unavoidable race condition that occurs when checking if the database file
// exists.
func LoadDatabase(name string) (*Database, error) {
	d := &Database{
		name:  name,
		Paths: make(map[string]string),
	}
	if _, err := os.Stat(name); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		r, err := os.Open(name)
		if err != nil {
			return nil, err
		}
		defer r.Close()
		if err := json.NewDecoder(r).Decode(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}

// Save attempts to write the database to disk. This should only need to be
// called when the contents of the map are changed.
func (d *Database) Save() error {
	if err := os.MkdirAll(path.Dir(d.name), 0775); err != nil {
		return err
	}
	w, err := os.Create(d.name)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := json.NewEncoder(w).Encode(d); err != nil {
		return err
	}
	return nil
}
