package errors

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(1, "success")
	log.Println(err.Code(), err.Error(), err.Message(), err.StackTrace())

	err = NewWithError(10000, err)
	log.Println(err.Code(), err.Error(), err.Message(), err.StackTrace())

}
