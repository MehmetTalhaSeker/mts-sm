package main

import (
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/database"
	"github.com/MehmetTalhaSeker/mts-sm/internal/fs"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/logg"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"github.com/MehmetTalhaSeker/mts-sm/route"
	"github.com/MehmetTalhaSeker/mts-sm/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"os"
)

func main() {
	conf := config.Init()
	logg.Init(conf)

	mc, err := fs.Init(conf)
	if err != nil {
		logg.L.Fatal("Minio connection failed: $s", zap.Error(err))
	}

	db, err := database.Init(conf)
	if err != nil {
		logg.L.Fatal("Database Connection Message $s", zap.Error(err))
	}

	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		logg.L.Fatal("Gorm Migration Failed: $s", zap.Error(err))
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorutils.ErrorHandler,
	})
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		// %s |%s %3d %s| %7v | %15s |%s %-7s %s
		Format:     "[${time}] | ${yellow}${status}${reset} | ${latency} 	| ${method} ${path} ${queryParams}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     os.Stderr,
	}))

	logg.L.Info("Application:\n",
		zap.String("environment", conf.Env),
		zap.String("port", conf.Rest.Port),
	)

	var (
		ur = repository.NewUserRepository(db)
	)

	var (
		as = service.NewAuthService(ur)
		us = service.NewUserService(ur, mc)
	)

	v1 := app.Group("v1")
	route.AuthRouter(v1, as)
	route.UserRouter(v1, as, us)

	logg.L.Fatal("Started", zap.Error(app.Listen(fmt.Sprintf(":%s", conf.Rest.Port))))
}
