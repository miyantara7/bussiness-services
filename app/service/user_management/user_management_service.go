package user_mangement

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	ucUser "github.com/vins7/bussiness-services/app/usecase/user_management"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/bussiness"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserManagementService struct {
	uc ucUser.UserManagement
}

func NewUserManagementService(uc ucUser.UserManagement) *UserManagementService {
	return &UserManagementService{
		uc: uc,
	}
}

func (u *UserManagementService) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	res, err := u.uc.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}

	var out *proto.LoginResponse
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}

func (u *UserManagementService) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*empty.Empty, error) {
	if err := u.uc.RegisterUser(ctx, req); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (u *UserManagementService) GetDetailUserInformation(ctx context.Context, req *proto.UserInformationReq) (*proto.UserInformationRes, error) {
	res, err := u.uc.GetDetailUserInformation(ctx, req)
	if err != nil {
		return nil, err
	}

	var out *proto.UserInformationRes
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}
