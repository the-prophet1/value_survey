package server

import (
	. "value-survey/pkg/errors"
)

type Validator interface {
	Valid() Error
}

func validInterface(validator interface{}) Error {
	if validator, ok := validator.(Validator); ok {
		if err := validator.Valid(); err != nil {
			return err
		}
	}
	return nil
}
