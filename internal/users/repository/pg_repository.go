package repository

import (
	"database/sql"
	"log"

	"github.com/Aserold/go-crud/internal/models"
	"github.com/Aserold/go-crud/internal/users"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) users.Repository {
	return &usersRepo{db: db}
}

// GetByEmail implements users.Repository.
func (r *usersRepo) FindByEmail(user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	if err := r.db.QueryRowx(findUserByEmail, user.Email).StructScan(foundUser); err != nil {
		return nil, errors.Wrap(err, "usersRepo.FindByEmail.QueryRowx")
	}
	return foundUser, nil
}

// Create implements users.Repository.
func (r *usersRepo) Create(user *models.User) (*models.User, error) {
	u := &models.User{}

	if err := r.db.QueryRowx(createUserQuery, &user.Username,
		&user.Email, &user.Age).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "usersRepo.Create.StructScan")
	}

	return u, nil
}

// Delete implements users.Repository.
func (r *usersRepo) Delete(userID int) error {
	result, err := r.db.Exec(deleteUserQuery, userID)

	if err != nil {
		return errors.WithMessage(err, "userRepo Delete Exec")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "userRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "userRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID implements users.Repository.
func (r *usersRepo) GetByID(userID int) (*models.User, error) {
	user := &models.User{}

	if err := r.db.QueryRowx(getUserQuery, userID).StructScan(user); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByID.QueryRowx")
	}
	return user, nil
}

// ListUsers implements users.Repository.
func (r *usersRepo) ListUsers() (*models.UsersList, error) {
	var totalCount int
	if err := r.db.Get(&totalCount, getTotal); err != nil {
		log.Print(err)
		return nil, errors.Wrap(err, "userRepo.ListUsers.Get.totalCount")
	}
	
	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			Users:      make([]*models.User, 0),
			}, nil
		}
		
		var users = make([]*models.User, 0, 10)
		if err := r.db.Select(&users, listUserQuery); err != nil {
			log.Print(err)
			return nil, errors.Wrap(err, "userRepo.ListUsers.Select")
		}
		
		return &models.UsersList{
			TotalCount: totalCount,
			Users:      users,
			}, nil
		}
		
		// Update implements users.Repository.
		func (r *usersRepo) Update(user *models.User) (*models.User, error) {
			u := &models.User{}
			if err := r.db.Get(u, updateUserQuery, &user.Username, &user.Email, &user.Age, &user.ID); err != nil {
				log.Print(err)
				return nil, errors.Wrap(err, "userRepo.Update.Get")
			}
			
			return u, nil
		}
