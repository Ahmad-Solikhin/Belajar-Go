package Golang_Logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	log := logrus.New()

	log.Println("Hello Logger")
	fmt.Println("Hello Logger")
}

func TestLevel(t *testing.T) {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	log.Trace("Trace")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
}
