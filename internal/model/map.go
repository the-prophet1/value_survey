package model

import (
	"database/sql/driver"
	"encoding/json"
	"value-survey/pkg/errors"
)

type SIMap map[string]Int

func (m SIMap) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New(codeToByteSliceErr, "convert ot byte slice error")
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	return nil
}

func (m SIMap) Value() (driver.Value, error) {
	if m == nil {
		return []byte{}, nil
	}
	return json.Marshal(m)
}
