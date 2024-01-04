package configs

import "os"

var (
	dbcfg   *DBConfig
	DB_HOST = "DB_HOST"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func GetDB() *DBConfig {
	dbcfg = &DBConfig{
		Host:     os.Getenv(DB_HOST),
		Port:     os.Getenv(DB_PORT),
		User:     os.Getenv(DB_USER),
		Pass:     os.Getenv(DB_PASS),
		Database: os.Getenv(DB_NAME),
	}

	return dbcfg
}
