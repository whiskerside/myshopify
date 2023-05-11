package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/whiskerside/myshopify/conf"
)

var (
	logger  zerolog.Logger
	logFile *os.File
	err     error
)

func init() {
	zerolog.TimestampFieldName = "T"
	zerolog.LevelFieldName = "L"
	zerolog.MessageFieldName = "M"

	logFile, err = os.OpenFile(conf.Env.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	logger = zerolog.New(logFile).Level(levelConf(conf.Env.Log.Level))
}

func Logger() zerolog.Logger {
	return logger.With().Timestamp().Caller().Logger()
}

func levelConf(l string) zerolog.Level {
	switch l {
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.DebugLevel
	}
}
