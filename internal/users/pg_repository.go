//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package users

import "github.com/Aserold/go-crud/internal/models"

type Repository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID int) error
	GetByID(userID int) (*models.User, error)
	FindByEmail(user *models.User) (*models.User, error)
	ListUsers() (*models.UsersList, error)
}