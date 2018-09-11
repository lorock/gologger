package main

import (
	"github.com/lorock/gologger"
)

func main() {

	logger := gologger.NewLogger()

	fileConfig := &gologger.FileConfig{
		Filename: "./logs/test.log",
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): "./logs/error.log",
			logger.LoggerLevel("info"):  "./logs/info.log",
			logger.LoggerLevel("debug"): "./logs/debug.log",
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

		i += 1
		if i == 21000 {
			break
		}
	}

	logger.Flush()
}