package Golang_Logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHookWarn struct {
}

func (s *SampleHookWarn) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel}
}

func (s *SampleHookWarn) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook Warn", entry.Level, entry.Message)
	return nil
}

type SampleHookError struct {
}

func (s *SampleHookError) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (s *SampleHookError) Fire(entry *logrus.Entry) error {
	fmt.Printf("Sample Hook Error : %s, %s\n", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	log := logrus.New()
	log.AddHook(&SampleHookWarn{})
	log.AddHook(&SampleHookError{})

	log.Info("Hello Ges")
	log.Warn("Nah ini Warn")
	log.Error("Nah ini Error")
}
