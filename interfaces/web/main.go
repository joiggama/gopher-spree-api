package main

import (
	"os"

	_ "github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/interfaces/web/api"
	"github.com/crowdint/gopher-spree-api/logger"
    _ "github.com/jpfuentes2/go-env/autoload"
)

func main() {
	err := repositories.InitDB(false)
	logrus := logger.NewLogrus()
	logger.SetLogger(logrus)
	logger.SetLevel(configs.Get(configs.LOG_LEVEL))

	if err != nil {
		panic(err)
	}

	api.Router().Run("0.0.0.0:" + os.Getenv("PORT"))
}
