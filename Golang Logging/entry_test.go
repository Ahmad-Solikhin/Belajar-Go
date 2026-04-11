package Golang_Logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(log)

	entry.Info("Haiyaaa Looo")
}
