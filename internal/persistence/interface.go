package persistence

import (
	. "value-survey/pkg/errors"
)

type BalanceSheet struct {
	ID int
}

type Interface interface {
	CreateBalanceSheet(sheet *BalanceSheet) Error
}
