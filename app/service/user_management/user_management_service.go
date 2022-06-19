package user_mangement

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	ucUser "github.com/vins7/bussiness-services/app/usecase/user_management"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
)

type UserManagementService struct {
	uc ucUser.UserManagement
}

func NewUserManagementService(uc ucUser.UserManagement) *UserManagementService {
	return &UserManagementService{
		uc: uc,
	}
}

func (u *UserManagementService) Login(ctx context.Context, req *proto.LoginRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}
