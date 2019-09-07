package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
golang标准库的日志框架非常简单，仅仅提供了print,panic和fatal三个函数,对于更精细
的日志级别，日志文件分割以及日志分发等方面并没有提供支持
logrus完全兼容golang标准库日志模块
*/

var log = logrus.New()

func main() {
	// 输出设置为标准输出
	log.Out = os.Stdout

	// 消息输出格式为json格式
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// {"animal":"walrus","level":"info","msg":"A group of walrus emerges from the ocean","size":10,"time":"2019-07-29T12:19:27+08:00"}
}
