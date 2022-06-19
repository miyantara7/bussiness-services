package user_management

import (
	repo "github.com/vins7/bussiness-services/app/adapter/client/user_management"
)

type UserManagementUsecase struct {
	repo repo.UserManagementRepo
}

func NewUserManagementUsecase(repo repo.UserManagementRepo) UserManagement {
	return &UserManagementUsecase{
		repo: repo,
	}
}

func (u *UserManagementUsecase) Login() {

}
