package handler

import (
	"value-survey/internal/model"
	"value-survey/internal/persistence"
	. "value-survey/pkg/errors"
)

type BalanceSheetHandler struct {
}

func newBalanceSheetHandler() *BalanceSheetHandler {
	return &BalanceSheetHandler{}
}

func GetBalanceSheetHandler() *BalanceSheetHandler {
	return balanceSheetHandler
}

type BalanceSheetCreateReq struct {
	Ticket int // 股票号码
	Name   string
	model.CurrentAsset
	model.NonCurrentAsset
	model.CurrentLiability
	model.NonCurrentLiability
}

type BalanceSheetCreateResp struct {
}

func (b *BalanceSheetHandler) Create(req *BalanceSheetCreateReq) (*BalanceSheetCreateResp, Error) {
	balanceSheet := &persistence.BalanceSheet{
		Ticket: req.Ticket,
		Name:   req.Name,
	}
	balanceSheet.ContractAsset = req.ContractAsset
	balanceSheet.NonCurrentAsset = req.NonCurrentAsset
	balanceSheet.CurrentLiability = req.CurrentLiability
	balanceSheet.NonCurrentLiability = req.NonCurrentLiability
	if err := b.storage.CreateBalanceSheet(balanceSheet); err != nil {
		return nil, err
	}
	return &BalanceSheetCreateResp{}, nil
}
