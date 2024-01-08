package logg

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"go.uber.org/zap"
	"log"
)

var L *zap.Logger

func Init(conf *config.Config) {
	var err error
	if conf.Env == "production" {
		L, err = zap.NewProduction()

		if err != nil {
			log.Panic(err)
		}
	} else {
		L, err = zap.NewDevelopment()
		if err != nil {
			log.Panic(err)
		}
	}

	L.Info("Log initialize")

}
