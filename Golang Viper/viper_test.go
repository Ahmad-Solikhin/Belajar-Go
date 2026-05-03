package Golang_Viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("app.name"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("app.name"))
}

func TestEnv(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))
}

func TestEnvVariable(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))

	assert.Equal(t, "hello", config.GetString("FROM_ENV"))
}
