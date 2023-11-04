package config

import (
	"github.com/hobord/poc-htmx-go-todolist/dal/config"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

type service struct {
	reader config.Reader
}

func NewService(reader config.Reader) Service {
	return &service{
		reader: reader,
	}
}

func (s *service) GetServerConfig() (entities.ServerConfig, error) {
	cfg, _ := s.reader.ReadServerConfig()

	err := s.validateServerConfig(&cfg)

	return cfg, err
}

func (s *service) validateServerConfig(cfg *entities.ServerConfig) error {
	if cfg.HttpPort == 0 {
		cfg.HttpPort = DefaultHttpPort
	}

	return nil
}
