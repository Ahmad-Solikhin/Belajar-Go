package Golang_Logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	log := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)

	log.Info("Hayolo")
	log.Warn("Hayolo Warn")
	log.Error("Hayolo Error")
}
