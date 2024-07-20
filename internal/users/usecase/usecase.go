package usecase

import (
	"net/http"

	"github.com/Aserold/go-crud/config"
	"github.com/Aserold/go-crud/internal/models"
	"github.com/Aserold/go-crud/internal/users"
	"github.com/Aserold/go-crud/pkg/httpErrors"
)

type usersUC struct {
	cfg       *config.Config
	usersRepo users.Repository
}

// Create implements users.UseCase.
func (u *usersUC) Create(user *models.User) (*models.User, error) {
	existsUser, err := u.usersRepo.FindByEmail(user)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}

	createdUser, err := u.usersRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

// Delete implements users.UseCase.
func (u *usersUC) Delete(userID int) error {
	return u.usersRepo.Delete(userID)
}

// GetByID implements users.UseCase.
func (u *usersUC) GetByID(userID int) (*models.User, error) {
	return u.usersRepo.GetByID(userID)
}

// ListUsers implements users.UseCase.
func (u *usersUC) ListUsers() (*models.UsersList, error) {
	return u.usersRepo.ListUsers()
}

// Update implements users.UseCase.
func (u *usersUC) Update(user *models.User) (*models.User, error) {
	return u.usersRepo.Update(user)
}

func NewUsersUseCase(cfg *config.Config, usersRepo users.Repository) users.UseCase {
	return &usersUC{cfg: cfg, usersRepo: usersRepo}
}
