package gotool

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestElasticHook(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL(""),
		elastic.SetBasicAuth("", ""),
		elastic.SetSniff(false))
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	hook, err := NewElasticHook(client, "localhost", logrus.WarnLevel, "testlog", RotateDay)
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	logrus.AddHook(hook)

	logrus.WithFields(logrus.Fields{
		"name": "spotmax",
		"age":  15,
	}).Warn("Hello world!")

	logrus.WithFields(logrus.Fields{
		"name": "spotmax",
		"age":  100,
	}).Info("Hello world!")
}
