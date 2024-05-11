package persistence

import (
	"gorm.io/gorm"
	"value-survey/internal/model"
	. "value-survey/pkg/errors"
)

type BalanceSheet struct {
	ID     int    `json:"id" gorm:"column:id;primaryKey;comment:主键ID"`
	Ticket int    `json:"ticket" gorm:"column:ticket;comment:号码"`
	Name   string `json:"name" gorm:"column:name;comment:名称"`
	model.CurrentAsset
	model.NonCurrentAsset
	model.CurrentLiability
	model.NonCurrentLiability
}

type Interface interface {
	CreateBalanceSheet(sheet *BalanceSheet) Error
}

func InitRDB(db *gorm.DB) {
	rdbDefaultClient = rdbClient{db: db}
}
