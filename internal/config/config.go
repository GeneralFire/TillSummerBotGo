package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type TokenDescriptor struct {
	Token string `yaml:"APIToken"`
}

func GetTokenYAML(absPath string) (string, error) {
	rawYAML, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	var td TokenDescriptor
	err = yaml.Unmarshal(rawYAML, &td)
	if err != nil {
		return "", err
	}
	return td.Token, nil
}
