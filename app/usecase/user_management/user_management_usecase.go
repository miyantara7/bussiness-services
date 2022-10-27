package user_management

import (
	"context"

	repo "github.com/vins7/bussiness-services/app/adapter/client/user_management_services"
)

type UserManagementUsecase struct {
	repo repo.UserManagementRepo
}

func NewUserManagementUsecase(repo repo.UserManagementRepo) UserManagement {
	return &UserManagementUsecase{
		repo: repo,
	}
}

func (u *UserManagementUsecase) LoginUser(ctx context.Context, in interface{}) (interface{}, error) {
	u.repo.Login(in)
	return nil, nil
}

func (u *UserManagementUsecase) RegisterUser(ctx context.Context, in interface{}) error {
	u.repo.Register(in)
	return nil
}

func (u *UserManagementUsecase) GetDetailUserInformation(ctx context.Context, in interface{}) (interface{}, error) {
	u.repo.UserInformation(in)
	return nil, nil
}
