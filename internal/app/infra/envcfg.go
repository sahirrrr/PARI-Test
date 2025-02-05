package infra

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	DB    = "DB"
	ECHO  = "APP"
	REDIS = "REDIS"
)

func LoadDatabaseCfg() (*DatabaseCfg, error) {
	var cfg DatabaseCfg

	prefix := DB

	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}

// LoadEchoCfg load env to new instance of EchoCfg.
func LoadEchoCfg() (*AppCfg, error) {
	var cfg AppCfg

	prefix := ECHO

	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}
