package Golang_Logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	log := logrus.New()

	log.WithField("username", "Gayuh").Info("Hello World")

	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithField("username", "Gayuh").Info("Hello World")
}

func TestFields(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{"username": "asgr39", "name": "Gayuh"}).Info("Hello World")
}
