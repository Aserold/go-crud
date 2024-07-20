package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Aserold/go-crud/config"
	"github.com/Aserold/go-crud/internal/models"
	"github.com/Aserold/go-crud/internal/users"
	"github.com/Aserold/go-crud/pkg/httpErrors"
	"github.com/Aserold/go-crud/pkg/utils"
	"github.com/labstack/echo/v4"
)

type usersHandlers struct {
	cfg     *config.Config
	usersUC users.UseCase
}

// NewAuthHandlers Auth handlers constructor
func NewUsersHandlers(cfg *config.Config, usersUC users.UseCase) users.Handlers {
	return &usersHandlers{cfg: cfg, usersUC: usersUC}
}

// Create implements users.Handlers.
func (u *usersHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			log.Printf("bad request")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := u.usersUC.Create(user)
		if err != nil {
			log.Printf("couldn't create user")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

// Delete implements users.Handlers.
func (u *usersHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Println("couldn't parse id")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err := u.usersUC.Delete(uID); err != nil {
			log.Println("couldn't delete id")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID implements users.Handlers.
func (u *usersHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Println("couldn't parse id")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := u.usersUC.GetByID(uID)
		if err != nil {
			log.Println("couldn't get user by id")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, user)
	}
}

// ListUsers implements users.Handlers.
func (u *usersHandlers) ListUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		usersList, err := u.usersUC.ListUsers()
		if err != nil {
			log.Println("couldn't list users")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, usersList)
	}
}

// Update implements users.Handlers.
func (u *usersHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Println("couldn't parse id")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user := &models.User{}
		user.ID = int64(uID)

		if err = utils.ReadRequest(c, user); err != nil {
			log.Println("bad request")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedUser, err := u.usersUC.Update(user)
		if err != nil {
			log.Println("couldn't update user")
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedUser)
	}
}
