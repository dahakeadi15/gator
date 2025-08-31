package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding home dir: %w", err)
	}
	configFilePath := homeDirPath + "/" + configFileName

	return configFilePath, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error converting to bytes: %w", err)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	configFile, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer configFile.Close()

	_, err = configFile.Write(data)
	if err != nil {
		return fmt.Errorf("error writing data: %w", err)
	}
	return nil
}
