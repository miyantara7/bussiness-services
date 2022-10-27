package user_management_services

import (
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	cfgServer "github.com/vins7/bussiness-services/config"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserManagementClient struct {
	client proto.UsermanagementServiceClient
}

func NewUserManagementClient(c *grpc.ClientConn) UserManagementRepo {
	return &UserManagementClient{
		client: proto.NewUsermanagementServiceClient(c),
	}
}

func (u *UserManagementClient) Login(in interface{}) (*proto.LoginResponse, error) {
	data, ok := in.(*model.LoginRequest)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}
	res, err := u.client.Login(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.LoginRequest{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserManagementClient) Register(in interface{}) error {
	data, ok := in.(*model.CreateUserReq)
	if !ok {
		return status.Errorf(codes.Internal, "Error while casting")
	}
	_, err := u.client.Register(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.RegisterRequest{
		Username:    data.Username,
		Password:    data.Password,
		Nama:        data.Nama,
		Email:       data.Email,
		NoIdentitas: data.NoIdentitas,
		TglLahir:    data.TglLahir,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserManagementClient) UserInformation(in interface{}) (*proto.UserInformationRes, error) {
	data, ok := in.(*model.DetailUserReq)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}
	res, err := u.client.UserInformation(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.UserInformationReq{
		UserID:   data.UserId,
		Username: data.Username,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
