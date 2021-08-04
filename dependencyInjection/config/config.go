package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "root:123456@tcp(127.0.0.1:3306)/forge",
		Port:         "8000",
	}
}

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("mysql", config.DatabasePath)
}
