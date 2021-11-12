package config

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port                 string `yaml:"port"`
	CacheExpirationTime  int    `yaml:"cache_expiration_time"`
	CacheCleanupInterval int    `yaml:"cache_cleanup_interval"`
}

func ReadConfig(configFile string) (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("error while opening config file: %w", err)
	}

	var cfg Config
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("error while decoding config from yaml: %w", err)
	}

	if err = cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("error while validating config: %w", err)
	}
	return cfg, nil
}

func (c Config) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, is.Port),
		validation.Field(&c.CacheExpirationTime, validation.Required),
		validation.Field(&c.CacheCleanupInterval, validation.Required),
	)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
