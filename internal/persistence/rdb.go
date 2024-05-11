package persistence

import (
	"gorm.io/gorm"
	. "value-survey/pkg/errors"
)

const (
	codeOpenSQLiteErr = iota + 20000
)

var rdbDefaultClient rdbClient

type rdbClient struct {
	db *gorm.DB
}

func (s *rdbClient) CreateBalanceSheet(sheet *BalanceSheet) Error {

	return nil
}
