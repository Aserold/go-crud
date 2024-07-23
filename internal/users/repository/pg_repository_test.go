package repository

import (
	"fmt"
	"testing"

	"github.com/Aserold/go-crud/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUsersRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"username", "email", "age"}).AddRow(
			"aserold",
			"aserold@gmail.com",
			uint8(18),
		)

		user := &models.User{
			Username: "aserold",
			Email:    "aserold@gmail.com",
			Age:      18,
		}

		mock.ExpectQuery(createUserQuery).WithArgs(
			&user.Username,
			&user.Email,
			&user.Age,
		).WillReturnRows(rows)
		createdUser, err := usersRepo.Create(user)

		require.NoError(t, err)
		require.NotNil(t, createdUser)
		require.Equal(t, createdUser, user)
	})
}

func TestUsersRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"username", "email", "age"}).AddRow(
			"aserold",
			"aserold@gmail.com",
			uint8(18),
		)

		user := &models.User{
			Username: "aserold",
			Email:    "aserold@gmail.com",
			Age:      18,
		}

		mock.ExpectQuery(updateUserQuery).WithArgs(
			&user.Username,
			&user.Email,
			&user.Age,
			&user.ID,
		).WillReturnRows(rows)
		updatedUser, err := usersRepo.Update(user)

		require.NoError(t, err)
		require.NotNil(t, updatedUser)
		require.Equal(t, updatedUser, user)
	})
}

func TestUsersRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		id := 1

		mock.ExpectExec(deleteUserQuery).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := usersRepo.Delete(id)
		require.Nil(t, err)
	})
}

func TestUsersRepo_GetByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("GetByID", func(t *testing.T) {
		id := 1

		rows := sqlmock.NewRows([]string{"id", "username", "email", "age"}).AddRow(
			id,
			"aserold",
			"aserold@gmail.com",
			uint8(18),
		)
		testUser := &models.User{
			ID:       1,
			Username: "aserold",
			Email:    "aserold@gmail.com",
			Age:      18,
		}

		mock.ExpectQuery(getUserQuery).WithArgs(id).WillReturnRows(rows)
		user, err := usersRepo.GetByID(id)
		require.NoError(t, err)
		require.Equal(t, user, testUser)
		fmt.Printf("test user: %s \n", testUser.Username)
		fmt.Printf("user: %s \n", user.Username)
	})
}

func TestUsersRepo_FindByEmail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("FindByEmail", func(t *testing.T) {
		id := 1

		rows := sqlmock.NewRows([]string{"id", "username", "email", "age"}).AddRow(
			id,
			"aserold",
			"aserold@gmail.com",
			uint8(18),
		)
		testUser := &models.User{
			ID:       1,
			Username: "aserold",
			Email:    "aserold@gmail.com",
			Age:      18,
		}

		mock.ExpectQuery(findUserByEmail).WithArgs(testUser.Email).WillReturnRows(rows)

		foundUser, err := usersRepo.FindByEmail(testUser)
		require.NoError(t, err)
		require.NotNil(t, foundUser)
		require.Equal(t, foundUser.Username, testUser.Username)
	})
}

func TestUsersRepo_ListUsers(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	usersRepo := NewUserRepository(sqlxDB)

	t.Run("ListUsers", func(t *testing.T) {
		id := 1

		totalCountRows := sqlmock.NewRows([]string{"count"}).AddRow(0)

		rows := sqlmock.NewRows([]string{"id", "username", "email", "age"}).AddRow(
			id,
			"aserold",
			"aserold@gmail.com",
			uint8(18),
		)

		mock.ExpectQuery(getTotal).WillReturnRows(totalCountRows)
		mock.ExpectQuery(listUserQuery).WillReturnRows(rows)

		users, err := usersRepo.ListUsers()
		require.NoError(t, err)
		require.NotNil(t, users)
	})
}
