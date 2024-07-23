package usecase

import (
	"database/sql"
	"testing"

	"github.com/Aserold/go-crud/config"
	"github.com/Aserold/go-crud/internal/models"
	"github.com/Aserold/go-crud/internal/users/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsersUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	mockUsersRepo := mock.NewMockRepository(ctrl)
	usersUC := NewUsersUseCase(cfg, mockUsersRepo)

	user := &models.User{
		Email: "aserold@gmail.com",
	}

	mockUsersRepo.EXPECT().FindByEmail(gomock.Eq(user)).Return(nil, sql.ErrNoRows)
	mockUsersRepo.EXPECT().Create(gomock.Eq(user)).Return(user, nil)

	createdUser, err := usersUC.Create(user)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.Nil(t, err)
}

func TestUsersUC_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	mockUsersRepo := mock.NewMockRepository(ctrl)
	usersUC := NewUsersUseCase(cfg, mockUsersRepo)

	user := &models.User{
		Username: "123456",
		Email:    "email@gmail.com",
	}

	mockUsersRepo.EXPECT().FindByEmail(gomock.Eq(user)).Return(nil, sql.ErrNoRows)
	mockUsersRepo.EXPECT().Update(gomock.Eq(user)).Return(user, nil)

	updatedUser, err := usersUC.Update(user)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	require.Nil(t, err)
}

func TestUsersUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	mockUsersRepo := mock.NewMockRepository(ctrl)
	usersUC := NewUsersUseCase(cfg, mockUsersRepo)

	user := &models.User{
	}

	mockUsersRepo.EXPECT().Delete(gomock.Eq(int(user.ID))).Return(nil)

	err := usersUC.Delete(int(user.ID))
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestUsersUC_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	mockUsersRepo := mock.NewMockRepository(ctrl)
	usersUC := NewUsersUseCase(cfg, mockUsersRepo)

	user := &models.User{
	}

	mockUsersRepo.EXPECT().GetByID(gomock.Eq(int(user.ID))).Return(user, nil)

	u, err := usersUC.GetByID(int(user.ID))
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, u)
}

func TestUsersUC_ListUsers(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	mockUsersRepo := mock.NewMockRepository(ctrl)
	usersUC := NewUsersUseCase(cfg, mockUsersRepo)

	usersList := &models.UsersList{}

	mockUsersRepo.EXPECT().ListUsers().Return(usersList, nil)

	users, err := usersUC.ListUsers()
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, users)
}
