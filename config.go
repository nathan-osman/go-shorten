package main

import (
	"encoding/json"
	"os"
)

// Config stores configuration data for the application. Part of the reason for
// requiring the information to come from a file instead of command-line args
// is because then the authentication data would not necessarily be safe from
// snooping.
type Config struct {
	Addr     string `json:"addr"`
	Database string `json:"database"`
	Admin    struct {
		Path     string `json:"path"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"admin"`
}

// LoadConfig attempts to load the configuration from a JSON file.
func LoadConfig(name string) (*Config, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
