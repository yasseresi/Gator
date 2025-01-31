package config

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
~/.gatorconfig.json and returns a Config struct. It should read the file from the HOME directory, then decode the JSON string into a new Config struct. I used os.UserHomeDir to get the location of HOME.
*/

func ReadConfig() Config {
	fileUrl, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error in getting file path")
	}
	if data, err := os.ReadFile(fileUrl); err != nil {
		fmt.Println("Error in reading file")
	} else {
		// Decode the JSON string into a Config struct
		var config Config
		json.Unmarshal(data, &config)
		return config
	}
	return Config{}

}
