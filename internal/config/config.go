package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"surlit/internal/logic/errors"
)

type Config struct {
	DBConnectionString string `yaml:"db_connection_string"`
}

func GetConfigENV() (*Config, error) {
	dbConnStr := os.Getenv("DB_CONNECTION_STRING")
	if len(dbConnStr) == 0 {
		return nil, errors.ErrCantGetDBConnectionString
	}
	return &Config{
		DBConnectionString: dbConnStr,
	}, nil
}

func GetConfigYML(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cant read yaml file: %w", err)
	}
	conf := Config{}
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, fmt.Errorf("cant unmarshal yaml file: %w", err)
	}
	return &conf, nil
}
