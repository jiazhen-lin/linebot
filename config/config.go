package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	LineConfig     LineBot  `yaml:"linebot"`
	DatabaseConfig Database `yaml:"database"`
	ServerConfig   Server   `yaml:"server"`
}

type LineBot struct {
	Secret string `yaml:"secret"`
	Token  string `yaml:"token"`
}

type Database struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Server struct {
	Port string `yaml:"port"`
}

type ConfigManager struct {
	cfg *Config
}

func (mgr *ConfigManager) loadFile() (*Config, error) {
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

func (mgr *ConfigManager) Get() (*Config, error) {
	if mgr.cfg != nil {
		return mgr.cfg, nil
	}
	c, err := mgr.loadFile()
	mgr.cfg = c
	if err != nil {
		return nil, err
	}
	return mgr.cfg, nil
}
