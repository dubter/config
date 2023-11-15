package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Port         string        `yaml:"port"`
	Host         string        `yaml:"host"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func LoadConfig(path string) (*YamlConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Failed to read httpconfig file: " + err.Error())
	}

	var config YamlConfig
	err = yaml.UnmarshalStrict(data, &config)
	if err != nil {
		return nil, errors.New("Failed to unmarshal httpconfig file: " + err.Error())
	}

	return &config, nil
}
