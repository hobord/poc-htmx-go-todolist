package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/hobord/poc-htmx-go-todolist/entities"
)

const (
	yamlFileType string = "yaml"
)

func (r *reader) ReadServerConfig() (entities.ServerConfig, error) {
	cfg := entities.ServerConfig{}

	v := viper.New()

	v.AddConfigPath(getConfigDirectory())
	v.SetConfigName(getConfigFileName())
	v.SetConfigType(yamlFileType)

	if err := v.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("could not read config. error %s", err)
	}

	var config ServerConfig
	if err := v.Unmarshal(&config); err != nil {
		return cfg, fmt.Errorf("could not unmarshal config. error %s", err)
	}

	applyDtoToEntity(config, &cfg)

	return cfg, nil
}

func getConfigDirectory() string {
	workingDirectory, _ := os.Getwd()

	return workingDirectory
}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	return fmt.Sprintf("config.%s", env)
}

func applyDtoToEntity(dto ServerConfig, cfg *entities.ServerConfig) {
	cfg.HttpPort = dto.HttpPort
}
