package main

import (
	"github.com/lorock/gologger"
)

func main() {

	logger := gologger.NewLogger()

	apiConfig := &gologger.ApiConfig{
		Url:        "http://127.0.0.1:8081/index.php",
		Method:     "GET",
		Headers:    map[string]string{},
		IsVerify:   false,
		VerifyCode: 0,
	}
	logger.Attach("api", gologger.LOGGER_LEVEL_DEBUG, apiConfig)
	logger.SetAsync()

	logger.Emergency("this is a emergency log!")
	logger.Alert("this is a alert log!")

	logger.Flush()
}
