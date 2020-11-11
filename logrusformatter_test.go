package gotool

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrusFormatter(t *testing.T) {
	LogrusFormat(logrus.DebugLevel)
	logrus.Debugf("%v","hello world!")
}