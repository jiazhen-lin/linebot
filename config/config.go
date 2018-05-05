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

// Manager is the manager that hold config instance
type Manager struct {
	cfg *Config
}

func (mgr *Manager) loadFile() (*Config, error) {
	cfg := &Config{}
	env := os.Getenv("ENV")
	var configPath string
	if env == "" {
		return nil, errors.New("Please specify ENV: PRODUCTION, DEVELOP, TEST")
	} else if env == "PRODUCTION" {
		configPath = "production.yml"
	} else if env == "DEVELOP" {
		configPath = "develop.yml"
	} else if env == "TEST" {
		configPath = "test.yml"
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

// Get is the function for returning config instance from config.Manager
func (mgr *Manager) Get() (*Config, error) {
	// return config instance directly if already loaded
	if mgr.cfg != nil {
		return mgr.cfg, nil
	}
	c, err := mgr.loadFile()
	if err != nil {
		return nil, err
	}
	mgr.cfg = c
	return mgr.cfg, nil
}
