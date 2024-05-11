package model

import (
	"database/sql/driver"
	"value-survey/pkg/errors"
)

const (
	codeToIntErr = 50000 + iota
	codeToByteSliceErr
)

type Int = *innerInt

type innerInt struct {
	ID         int            `json:"id" gorm:"column:id;primaryKey;comment:主键ID"`
	Data       int            `json:"data" gorm:"column:data;bigint;comment:数据"`
	Ingredient map[string]int `json:"ingredient,omitempty" gorm:"column:ingredient;charset(2048);comment:数据成分"`
}

func (i Int) TableName() string {
	return "inner_data"
}

func toInt(value interface{}) (int, error) {
	switch value := value.(type) {
	case int:
		return value, nil
	case int16:
		return int(value), nil
	case int8:
		return int(value), nil
	case int32:
		return int(value), nil
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
	i.ID = id

	return nil
}

func (i Int) Value() (driver.Value, error) {
	return i.Data, nil
}
