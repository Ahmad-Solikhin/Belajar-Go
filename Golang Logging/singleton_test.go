package Golang_Logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Error("Haiyaa Error")
}
