package gotool

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrusFormatter(t *testing.T) {
	LogrusFormat(logrus.DebugLevel)
	logrus.Debugf("%v", "hello world!")
}

func TestLogrusES(t *testing.T) {
	config := ESConfig{
		Url:       "",
		UserName:  "",
		Password:  "",
		IndexName: "testlog",
	}
	LogrusES(&config)
	logrus.WithFields(logrus.Fields{
		"name": "spotmax",
		"age":  15,
	}).Warn("Hello world!")

	logrus.WithFields(logrus.Fields{
		"name": "spotmax",
		"age":  100,
	}).Info("Hello world!")
}
