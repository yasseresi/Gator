package config

import "os"

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeUrl, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := homeUrl + "/" + configFileName
	return path, nil
}
