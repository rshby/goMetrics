package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GenerateConsoleLogging() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	return log
}
