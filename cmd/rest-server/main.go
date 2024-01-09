package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"

	"github.com/MehmetTalhaSeker/mts-sm/internal/database"
	"github.com/MehmetTalhaSeker/mts-sm/internal/fs"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/logg"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"github.com/MehmetTalhaSeker/mts-sm/route"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func main() {
	conf := config.Init()
	logg.Init(conf)

	mc, err := fs.New(conf)
	if err != nil {
		logg.L.Fatal("Minio connection failed: $s", zap.Error(err))
	}

	db, err := database.Init(conf)
	if err != nil {
		logg.L.Fatal("Database Connection Message $s", zap.Error(err))
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.PostLike{},
		&model.CommentLike{},
		&model.Friendship{},
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

	prometheus := fiberprometheus.New("mts-pro-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	logg.L.Info("Application:\n",
		zap.String("environment", conf.Env),
		zap.String("port", conf.Rest.Port),
	)

	var (
		ur  = repository.NewUserRepository(db)
		pr  = repository.NewPostRepository(db)
		pcr = repository.NewCommentRepository(db)
		lr  = repository.NewLikeRepository(db)
		fr  = repository.NewFriendshipRepository(db)
	)

	var (
		as  = service.NewAuthService(ur)
		us  = service.NewUserService(ur, mc, conf)
		ps  = service.NewPostService(pr, mc, conf)
		pcs = service.NewCommentService(pcr, conf)
		ls  = service.NewLikeService(lr, conf)
		frs = service.NewFriendshipService(fr, conf)
	)

	v1 := app.Group("v1")
	route.AuthRouter(v1, as)
	route.UserRouter(v1, as, us)
	route.PostRouter(v1, as, ps)
	route.CommentRouter(v1, as, pcs)
	route.LikeRouter(v1, as, ls)
	route.FriendshipRouter(v1, as, frs)

	logg.L.Fatal("Started", zap.Error(app.Listen(fmt.Sprintf(":%s", conf.Rest.Port))))
}
