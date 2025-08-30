package env

import (
	"errors"
	"os"
	config "wb/internal/config"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(dsnEnvName)

	if len(dsnEnvName) == 0 {
		return nil, errors.New("pg dsn not found")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
