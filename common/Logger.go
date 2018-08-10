package common

import (
	"github.com/alecthomas/log4go"
)

func LoggerConfiguration(filepath string) {
	//	l4g.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())             //输出到控制台,级别为DEBUG
	//l4g.AddFilter("file", l4g.DEBUG, l4g.NewFileLogWriter("test.log", false)) //输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件

	log4go.LoadConfiguration(filepath) //使用加载配置文件,类似与java的log4j.propertites

	//defer log4go.Close()                                                         //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件
}

func Debug(arg0 interface{}, args ...interface{}) {
	log4go.Debug(arg0, args)
}

func Info(arg0 interface{}, args ...interface{}) {
	log4go.Info(arg0, args)
}

func Error(arg0 interface{}, args ...interface{}) {
	log4go.Error(arg0, args)
}
