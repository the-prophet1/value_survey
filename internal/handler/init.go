package handler

import (
	. "value-survey/pkg/errors"
)

var (
	balanceSheetHandler *BalanceSheetHandler
)

func InitHandler() Error {
	if err != nil {
		return err
	}

	balanceSheetHandler = newBalanceSheetHandler()

	return nil
}
