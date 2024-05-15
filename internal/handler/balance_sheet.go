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
	Ticket                    int    `json:"ticket"` // 股票号码
	Name                      string `json:"name"`
	model.CurrentAsset        `json:",inline"`
	model.NonCurrentAsset     `json:",inline"`
	model.CurrentLiability    `json:",inline"`
	model.NonCurrentLiability `json:",inline"`
}

type BalanceSheetCreateResp struct {
}

func (b *BalanceSheetHandler) Create(req *BalanceSheetCreateReq) (*BalanceSheetCreateResp, Error) {
	balanceSheet := &persistence.BalanceSheet{
		Ticket: req.Ticket,
		Name:   req.Name,
	}
	balanceSheet.CurrentAsset = req.CurrentAsset
	balanceSheet.NonCurrentAsset = req.NonCurrentAsset
	balanceSheet.CurrentLiability = req.CurrentLiability
	balanceSheet.NonCurrentLiability = req.NonCurrentLiability
	if err := persistence.GetRDB().CreateBalanceSheet(balanceSheet); err != nil {
		return nil, err
	}
	return &BalanceSheetCreateResp{}, nil
}
