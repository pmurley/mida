package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Log = logrus.New()
var Lum *lumberjack.Logger = nil

func InitGlobalLogger(logfile string) {
	fileFormatter := new(logrus.TextFormatter)
	fileFormatter.FullTimestamp = true
	fileFormatter.DisableColors = true

	Lum = &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    50, //megabytes
		MaxBackups: 3,
		MaxAge:     180, //days
	}

	mWriter := io.MultiWriter(os.Stdout, Lum)
	Log.SetOutput(mWriter)

	consoleFormatter := new(logrus.TextFormatter)
	consoleFormatter.FullTimestamp = true
	consoleFormatter.ForceColors = true

	Log.SetOutput(os.Stdout)
	Log.SetFormatter(consoleFormatter)
}

func CloseGlobalLoggerFile() {
	err := Lum.Close()
	if err != nil {
		Log.Fatal(err)
	}
}
