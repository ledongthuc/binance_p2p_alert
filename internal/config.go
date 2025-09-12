package internal

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config holds all configuration values for the application
type Config struct {
	// Trading pair configuration
	Asset string `yaml:"asset"`
	Fiat  string `yaml:"fiat"`

	// Alert conditions
	MaxPrice      float64 `yaml:"max_price"`
	MinOfMaxAmount  float64 `yaml:"min_of_max_amount"`

	// API configuration
	PageSize int `yaml:"page_size"`

	EnableSlackAlert bool `yaml:"enable_slack_alert"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %v", filename, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %v", filename, err)
	}

	return &config, nil
}
