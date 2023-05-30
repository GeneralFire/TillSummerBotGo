package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type BotConfig struct {
	Token string `yaml:"APIToken"`
	DB    struct {
		Driver        string `yaml:"driver"`
		ConnectString string `yaml:"connect"`
	} `yaml:"SQL"`
}

func GetConfigYAML(absPath string) (*BotConfig, error) {
	rawYAML, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var config BotConfig
	err = yaml.Unmarshal(rawYAML, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
