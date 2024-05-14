package handler

import (
	. "value-survey/pkg/errors"
)

var (
	balanceSheetHandler *BalanceSheetHandler
)

func InitHandler() Error {
	balanceSheetHandler = newBalanceSheetHandler()
	return nil
}
