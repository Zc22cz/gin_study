package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
	Writer *os.File
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	hook.Writer.Write([]byte(line))
	return nil
}

func (hook *MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logrus.SetReportCaller(true)
	file, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	hook := &MyHook{Writer: file}
	logrus.AddHook(hook)
	logrus.Errorf("hello")
}
