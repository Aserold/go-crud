package server

import (
	usersHttp "github.com/Aserold/go-crud/internal/users/delivery/http"
	usersRepository "github.com/Aserold/go-crud/internal/users/repository"
	usersUseCase "github.com/Aserold/go-crud/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s Server) MapHandlers(e *echo.Echo) error {
	uRepo := usersRepository.NewUserRepository(s.db)

	usersUC := usersUseCase.NewUsersUseCase(s.cfg, uRepo)

	usersHandlers := usersHttp.NewUsersHandlers(s.cfg, usersUC)

	e.Use(middleware.Logger())
	e.Use(middleware.Secure())

	v1 := e.Group("/api/v1")

	usersGroup := v1.Group("/users")

	usersHttp.MapUsersRoutes(usersGroup, usersHandlers)

	return nil
}