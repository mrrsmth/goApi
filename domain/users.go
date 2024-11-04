package domain

import (
	"go_api_template/pkg/handlers"
	"go_api_template/pkg/repositories"
	"go_api_template/pkg/services"
	"go_api_template/pkg/utils"
)

func UserDomain(logger utils.ILogger, config utils.IConfig, db utils.PgPool) *handlers.UserHandler {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	logger.Success("Bootstrapped user domain.")

	return userHandler
}
