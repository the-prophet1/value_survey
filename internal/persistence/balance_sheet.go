package persistence

import "value-survey/internal/model"

type BalanceSheet struct {
	ID     int    `json:"id" gorm:"column:id;primaryKey;comment:主键ID"`
	Ticket int    `json:"ticket" gorm:"column:ticket;comment:号码"`
	Name   string `json:"name" gorm:"column:name;comment:名称"`
	model.CurrentAsset
	model.NonCurrentAsset
	model.CurrentLiability
	model.NonCurrentLiability
}

func (b *BalanceSheet) TableName() string {
	return "balance_sheet"
}
