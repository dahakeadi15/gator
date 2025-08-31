package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %w", err)
	}

	cfg := Config{}
	err = json.Unmarshal(configData, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error reading data: %w", err)
	}

	return cfg, nil
}
