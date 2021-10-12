package route

import (
	"Intern-project/pkg/handlers"
	"Intern-project/pkg/repo"
	"Intern-project/pkg/services"

	"github.com/caarlos0/env"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

func NewService() *Service {
	s := &Service{
		service.NewApp("Crawl data Service", "v1.0"),
		&extraSetting{},
	}
	_ = env.Parse(s.setting)

	db := s.GetDB()

	if s.setting.DbDebugEnable {
		db = db.Debug()
	}

	repo := repo.NewReop(db)

	userService := services.NewUserService(repo)
	authService := services.NewAuthService()

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService, authService)

	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/user/sign-up", ginext.WrapHandler(userHandler.SignUp))
	v1Api.POST("/auth/login", ginext.WrapHandler(authHandler.Login))

	// Metadata
	migrate := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrate.Migrate)
	return s
}
