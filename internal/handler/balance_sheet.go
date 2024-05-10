package handler

import (
	"value-survey/internal/model"
	. "value-survey/pkg/errors"
)

type BalanceSheetHandler struct {
}

func NewBalanceSheetHandler() *BalanceSheetHandler {
	return &BalanceSheetHandler{}
}

type BalanceSheetCreateReq struct {
	Ticket int // 股票号码
	model.CurrentAsset
	model.NonCurrentAsset
	model.CurrentLiability
	model.NonCurrentLiability
}

type BalanceSheetCreateResp struct {
}

func (b *BalanceSheetHandler) Create(req *BalanceSheetCreateReq) (*BalanceSheetCreateResp, Error) {
	return nil, nil
}
