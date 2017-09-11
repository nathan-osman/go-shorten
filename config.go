package shorten

import (
	"encoding/json"
	"os"
)

// Config stores configuration data for the application. Part of the reason for
// requiring the information to come from a file instead of command-line args
// is because then the authentication data would not necessarily be safe from
// snooping.
type Config struct {
	Addr          string `json:"addr"`
	Database      string `json:"database"`
	AdminPath     string `json:"admin_path"`
	AdminUsername string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
}

// WriteDefaultConfig writes the default configuration to a JSON file.
func WriteDefaultConfig(name string) error {
	c := &Config{
		Addr:          ":80",
		Database:      "db.json",
		AdminPath:     "/admin",
		AdminUsername: "admin",
		AdminPassword: "passw0rd",
	}
	w, err := os.Create(name)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := json.NewEncoder(w).Encode(c); err != nil {
		return err
	}
	return nil
}

// Load attempts to read the configuration from a JSON file.
func LoadConfig(name string) (*Config, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	c := &Config{}
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
