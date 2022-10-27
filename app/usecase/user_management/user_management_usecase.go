package user_management

import (
	"context"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/bussiness-services/app/adapter/client/user_management_services"
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	var req *model.LoginRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res, err := u.repo.Login(req)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Username: res.GetUsername(),
		UserId:   res.GetUserID(),
	}, nil
}

func (u *UserManagementUsecase) RegisterUser(ctx context.Context, in interface{}) error {
	var req *model.CreateUserReq
	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return u.repo.Register(req)
}

func (u *UserManagementUsecase) GetDetailUserInformation(ctx context.Context, in interface{}) (interface{}, error) {

	md, err := util.GetMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res, err := u.repo.UserInformation(&model.DetailUserReq{
		Username: md.UserName,
		UserId:   md.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &model.DataUser{
		Username:    res.GetUsername(),
		Nama:        res.GetNama(),
		UserId:      res.GetUserID(),
		Email:       res.GetEmail(),
		NoIdentitas: res.GetNoIdentitas(),
		TglLahir:    res.GetTglLahir(),
	}, nil
}
