package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"value-survey/internal/persistence"
	. "value-survey/pkg/errors"
)

const (
	codeOpenSQLiteErr = iota + 20000
)

type sqliteClient struct {
	db *gorm.DB
}

func (s *sqliteClient) CreateBalanceSheet(sheet *persistence.BalanceSheet) Error {
	return nil
}

func NewSQLite(path string) (persistence.Interface, Error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, NewWithError(codeOpenSQLiteErr, err)
	}
	return &sqliteClient{db: db}, nil
}
