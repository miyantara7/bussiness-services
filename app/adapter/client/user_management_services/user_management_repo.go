package user_management_services

import proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"

type UserManagementRepo interface {
	Login(in interface{}) (*proto.LoginResponse, error)
	Register(in interface{}) error
	UserInformation(in interface{}) (*proto.UserInformationRes, error)
}
