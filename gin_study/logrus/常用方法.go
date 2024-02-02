package main

import "github.com/sirupsen/logrus"

func main() {
	log1 := logrus.WithField("project", "study")
	log1.Errorln("hello")
	// time="2022-12-17T15:02:28+08:00" level=error msg=hello project=study
	log2 := logrus.WithFields(logrus.Fields{
		"func": "main",
	})
	log2.Warningf("你好")
	// time="2022-12-17T15:02:28+08:00" level=warning msg="你好" func=main
	log3 := log2.WithFields(logrus.Fields{
		"name": "杰",
	})
	// time="2022-12-17T15:02:28+08:00" level=warning msg="你好" name="杰" func=main
	log3.Warnln("你好")
}
