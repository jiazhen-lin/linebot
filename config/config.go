package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is top-level
type Config struct {
	LineConfig     LineBot  `yaml:"linebot"`
	DatabaseConfig Database `yaml:"database"`
	ServerConfig   Server   `yaml:"server"`
}

// LineBot is the field for LINE related config
type LineBot struct {
	Secret string `yaml:"secret"`
	Token  string `yaml:"token"`
}

// Database is the field for database related config
type Database struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Server is the field for server related config
type Server struct {
	Port string `yaml:"port"`
}

func loadFile() (*Config, error) {
	cfg := &Config{}
	env := os.Getenv("ENV")
	var configPath string
	if env == "" {
		return nil, errors.New("Please specify ENV: PRODUCTION, DEVELOP, TEST")
	} else if env == "PRODUCTION" {
		configPath = "config/production.yml"
	} else if env == "DEVELOP" {
		configPath = "config/develop.yml"
	} else if env == "TEST" {
		configPath = "config/test.yml"
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.UnmarshalStrict(content, cfg)
	if err != nil {
		return nil, fmt.Errorf("parsing YAML file %s: %v", configPath, err)
	}
	return cfg, nil
}

// New return a config from config file
func New() (*Config, error) {
	config, err := loadFile()
	if err != nil {
		return nil, err
	}
	return config, nil
}
