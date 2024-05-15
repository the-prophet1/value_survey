package model

import (
	"database/sql/driver"
	"encoding/json"
	"value-survey/pkg/errors"
	"value-survey/pkg/storage"
)

const (
	codeToIntErr = 50000 + iota
	codeToByteSliceErr
)

type innerInt struct {
	ID         int   `json:"id" gorm:"column:id;primaryKey;comment:主键ID"`
	Data       int   `json:"data" gorm:"column:data;bigint;comment:数据"`
	Ingredient siMap `json:"ingredient,omitempty" gorm:"column:ingredient;charset(2048);comment:数据成分"`
}

func (i Int) TableName() string {
	return "inner_data"
}

type Int = *innerInt

func NewInt() Int {
	return &innerInt{
		Ingredient: make(map[string]int),
	}
}

func toInt(value interface{}) (int, error) {
	switch value := value.(type) {
	case int64:
		return int(value), nil
	}
	return 0, errors.New(codeToIntErr, "convert to int error")
}

func (i Int) Scan(value interface{}) error {
	id, err := toInt(value)
	if err != nil {
		return err
	}
	if err := storage.GetDefault().Model(&innerInt{}).Where("id = ?", id).Find(i).Error; err != nil {
		return err
	}
	return nil
}

func (i Int) Value() (driver.Value, error) {
	if i == nil || i.Data == 0 {
		return nil, nil
	}
	return int64(i.ID), nil
}

type siMap map[string]int

func (m siMap) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New(codeToByteSliceErr, "convert ot byte slice error")
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	return nil
}

func (m siMap) Value() (driver.Value, error) {
	if m == nil {
		return []byte{}, nil
	}
	return json.Marshal(m)
}
