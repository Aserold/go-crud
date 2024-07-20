package users

import "github.com/Aserold/go-crud/internal/models"

type UseCase interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID int) error
	GetByID(userID int) (*models.User, error)
	ListUsers() (*models.UsersList, error)
}

