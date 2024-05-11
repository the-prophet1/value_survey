package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	. "value-survey/pkg/errors"
	"value-survey/pkg/storage"
)

const (
	codeTooManyPathErr = 40000 + iota
	codeOpenFileErr
	codeReadFileErr
	codeUnmarshalErr
)

var globalConfig Config

type Config struct {
	Persistence *storage.Config
}

func LoadConfig(path ...string) Error {
	if len(path) >= 2 {
		return New(codeTooManyPathErr, "too many paths")
	}
	var (
		file *os.File
		err  error
	)
	if len(path) == 0 {
		file, err = os.Open("./config/config.yaml")
	} else {
		file, err = os.Open(path[0])
	}
	if err != nil {
		return NewWithError(codeOpenFileErr, err)
	}

	configData, err := io.ReadAll(file)
	if err != nil {
		return NewWithError(codeReadFileErr, err)
	}

	err = yaml.Unmarshal(configData, &globalConfig)
	if err != nil {
		return NewWithError(codeUnmarshalErr, err)
	}
	return nil
}

func GetConfig() *Config {
	return &globalConfig
}
