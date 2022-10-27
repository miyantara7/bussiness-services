package user_mangement

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	ucUser "github.com/vins7/bussiness-services/app/usecase/user_management"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/bussiness"
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
	u.uc.LoginUser(ctx, req)
	return nil, nil
}

func (u *UserManagementService) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*empty.Empty, error) {
	u.uc.RegisterUser(ctx, req)
	return &empty.Empty{}, nil
}

func (u *UserManagementService) GetDetailUserInformation(ctx context.Context, req *proto.UserInformationReq) (*proto.UserInformationRes, error) {
	u.uc.GetDetailUserInformation(ctx, req)
	return nil, nil
}
