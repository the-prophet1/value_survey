package persistence

import (
	"gorm.io/gorm"
	"value-survey/internal/model"
	. "value-survey/pkg/errors"
)

const (
	codeAutoMigrateErr = iota + 20000
	codeCreateError
)

type Interface interface {
	CreateBalanceSheet(sheet *BalanceSheet) Error
}

func InitRDB(db *gorm.DB) Error {
	if err := db.AutoMigrate(&BalanceSheet{}, model.NewInt()); err != nil {
		return NewWithError(codeAutoMigrateErr, err)
	}
	rdbDefaultClient = rdbClient{db: db}
	return nil
}

func GetRDB() Interface {
	return &rdbDefaultClient
}
