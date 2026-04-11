package Golang_Logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestJsonFormatter(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("Test JSON Info")
}
