package config_test

import (
	"github.com/eneskzlcn/morsencoder/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	environment := "local"
	config, err := config.LoadConfig[config.Config]("../../.dev/", environment, "yaml")
	assert.Nil(t, err)
	assert.Equal(t, "4200", config.Server.Port)
}
