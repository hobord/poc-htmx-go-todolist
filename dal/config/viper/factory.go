package viper

import (
	"github.com/hobord/poc-htmx-go-todolist/dal/config"
)

type reader struct {
}

func NewReader() config.Reader {
	return &reader{}
}
