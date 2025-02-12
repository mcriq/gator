package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)


type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filepath.Join(homeDir, ".gatorconfig.json"))
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}