package persistence

import (
	"gorm.io/gorm"
	. "value-survey/pkg/errors"
)

var rdbDefaultClient rdbClient

type rdbClient struct {
	db *gorm.DB
}

func (s *rdbClient) CreateBalanceSheet(sheet *BalanceSheet) Error {
	sheet.CurrentAsset.GenerateID()
	sheet.NonCurrentAsset.GenerateID()
	sheet.CurrentLiability.GenerateID()
	sheet.NonCurrentLiability.GenerateID()
	if err := s.db.Model(&BalanceSheet{}).Create(sheet).Error; err != nil {
		return NewWithError(codeCreateError, err)
	}
	return nil
}
