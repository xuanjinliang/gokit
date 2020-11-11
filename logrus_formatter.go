package gotool

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func LogrusFormat(level logrus.Level) {
	FileLineFormatter := &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(FileLineFormatter)
}

type ESConfig struct {
	Url       string        `json:"url"`
	UserName  string        `json:"user_name"`
	Password  string        `json:"password"`
	Level     *logrus.Level `json:"level"`
	IndexName string        `json:"index_name"`
}

func LogrusES(config *ESConfig) {
	level := logrus.DebugLevel
	if config.Level != nil {
		level = *config.Level
	}
	LogrusFormat(level)
	client, err := elastic.NewClient(elastic.SetURL(config.Url),
		elastic.SetBasicAuth(config.UserName, config.Password),
		elastic.SetSniff(false))
	if err != nil {
		logrus.Panicf("%v", err.Error())
		return
	}
	host, err := os.Hostname()
	if err != nil {
		logrus.Panicf("%v", err.Error())
		return
	}
	hook, err := NewElasticHook(client, host, level, config.IndexName, RotateDay)
	if err != nil {
		logrus.Panicf("%v", err.Error())
		return
	}
	logrus.AddHook(hook)
}
