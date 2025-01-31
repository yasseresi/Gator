package config

import (
	"encoding/json"
	"os"
)

func Write(cfg Config) error {
	// Encode the Config struct into a JSON string
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	// Write the JSON string to the file
	fileUrl, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(fileUrl, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
