package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type ConfigFile struct {
	Config DBConfig `yaml:"DB"`
}

type DBConfig struct {
	DBPath string `yaml:"db_path"`
}

func GetConfig() (*DBConfig, error) {
	config := &ConfigFile{}
	path, err := filepath.Abs("../backend/config.yaml")
	if err != nil {
		return nil, err
	}

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, config)
	return &config.Config, err
}
