package storage

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DBName       string       `yaml:"DBName"`
	SQLiteConfig SQLiteConfig `yaml:"SQLiteConfig"`
}

func NewPersistence(config *Config) (*gorm.DB, error) {
	if config == nil {
		log.Warn("config is nil, use default db: SQLite")
		return newSQLite(defaultSqliteConfig)
	}

	switch config.DBName {
	case "sqlite":
		return newSQLite(config.SQLiteConfig)
	}

	log.Warnf("unknown db type: %s, use default db: SQLite", config.DBName)
	return newSQLite(defaultSqliteConfig)
}

var (
	defaultSqliteConfig = SQLiteConfig{Path: "default.db"}
)

type SQLiteConfig struct {
	Path string `yaml:"Path"`
}

var defaultDB *gorm.DB

func newSQLite(config SQLiteConfig) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
}

func InitDefault(config *Config) (err error) {
	defaultDB, err = NewPersistence(config)
	return
}

func GetDefault() *gorm.DB {
	return defaultDB
}
