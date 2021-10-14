// logger package wraps external logger (currently Logrus) to simplify future replacement
package logger

import (
	"github.com/sirupsen/logrus"
)

type LogLevel int

const (
	Verbose LogLevel = iota
	Debug
	Info
	Error
	Fatal
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:          true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func LogFatal(args ...interface{}) {
	logrus.Fatalln(args...)
}

func LogError(args ...interface{}) {
	logrus.Errorln(args...)
}

func LogInfo(args ...interface{}) {
	logrus.Infoln(args...)
}

func LogDebug(args ...interface{}) {
	logrus.Debugln(args...)
}

func LogVerbose(args ...interface{}) {
	logrus.Traceln(args...)
}

func SetLevel(lvl LogLevel) {
	switch lvl {
	case Verbose:
		logrus.SetLevel(logrus.TraceLevel)
	case Debug:
		logrus.SetLevel(logrus.DebugLevel)
	case Info:
		logrus.SetLevel(logrus.InfoLevel)
	case Error:
		logrus.SetLevel(logrus.InfoLevel)
	case Fatal:
		logrus.SetLevel(logrus.FatalLevel)
	default:
		LogError("Invalid log level set:", lvl)
	}
}
