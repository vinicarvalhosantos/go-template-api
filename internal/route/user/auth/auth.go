package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "gitlab.com/vinicius.csantos/go-template-api/internal/handler/user"
)

func SetupRegisterRoute(router fiber.Router) {
	register := router.Group("/register")

	//Register a new user
	register.Post("/", userHandler.RegisterUser)

}

func SetupLoginRoute(router fiber.Router) {
	login := router.Group("/login")

	//Sign in a user
	login.Post("/", userHandler.Login)

}
