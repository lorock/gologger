package main

import (
	"github.com/lorock/gologger"
)

var logdir string

func main() {

	logdir = "./logs/"
	logger := gologger.NewLogger()
	fileConfig := &gologger.FileConfig{
		Filename: logdir + "all.log",
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): logdir + "error.log",
			logger.LoggerLevel("info"):  logdir + "info.log",
			logger.LoggerLevel("debug"): logdir + "debug.log",
		},
		MaxSize:    1024 * 1024,
		MaxLine:    10000,
		DateSlice:  "d",
		JsonFormat: false,
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}
	logger.Attach("file", gologger.LOGGER_LEVEL_DEBUG, fileConfig)
	logger.SetAsync()

	i := 0
	for {
		logger.Emergency("this is a emergency log!")
		logger.Alert("this is a alert log!")
		logger.Critical("this is a critical log!")
		logger.Error("this is a error log!")
		logger.Warning("this is a warning log!")
		logger.Notice("this is a notice log!")
		logger.Info("this is a info log!")
		logger.Debug("this is a debug log!")

		logger.Emergency("this is a emergency log!")
		logger.Notice("this is a notice log!")
		logger.Info("this is a info log!")
		logger.Debug("this is a debug log!")

		logger.Emergency("this is a emergency log!")
		logger.Alert("this is a alert log!")
		logger.Critical("this is a critical log!")

		i++
		if i == 1050 {
			break
		}
	}

	logger.Flush()
}
