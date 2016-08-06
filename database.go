package main

import (
	"encoding/json"
	"os"
)

// Database maintains a list of subdomains and paths to use for redirects.
type Database struct {
	Subdomains map[string]string `json:"subdomains"`
	Paths      map[string]string `json:"paths"`
}

// LoadDatabase attempts to load the database from disk.
func LoadDatabase(name string) (*Database, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	d := &Database{}
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}

// Save attempts to write the database to disk. This should only need to be
// called when the contents of the map are changed.
func (d *Database) Save(name string) error {
	w, err := os.Create(name)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}
