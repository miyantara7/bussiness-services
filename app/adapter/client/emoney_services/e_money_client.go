package e_money_client

import (
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	cfgServer "github.com/vins7/bussiness-services/config"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/e_money_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EMoneyClient struct {
	client proto.BillerServiceClient
}

func NewEMoneyClient(c *grpc.ClientConn) EMoneyRepo {
	return &EMoneyClient{
		client: proto.NewBillerServiceClient(c),
	}
}

func (e *EMoneyClient) ListBiller() (*proto.BillerResponse, error) {

	res, err := e.client.ListBiller(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *EMoneyClient) DetailBiller(in interface{}) (*proto.Biller, error) {
	req, ok := in.(*model.BillerRequest)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}

	res, err := e.client.DetailBiller(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.BillerRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *EMoneyClient) GetBalance(in interface{}) (*proto.GetBalanceResponse, error) {
	req, ok := in.(*model.GetBalanceRequest)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}
	res, err := e.client.GetBalance(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.GetBalanceRequest{
			UserId:   req.UserId,
			UserName: req.UserName,
			NoKartu:  req.NoKartu,
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *EMoneyClient) CreateEMoney(in interface{}) (*proto.CreateEMoneyResponse, error) {
	req, ok := in.(*model.CreateEMoneyRequest)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}
	res, err := e.client.CreateEMoney(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.CreateEMoneyRequest{
			UserId:   req.UserId,
			UserName: req.UserName,
			Saldo:    req.Saldo,
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *EMoneyClient) GetTrxHist(in interface{}) (*proto.GetTrxHistResponse, error) {
	req, ok := in.(*model.GetTrxHistReq)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error while casting")
	}
	res, err := e.client.GetTrxHist(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.GetTrxHistReq{
			UserId:   req.UserId,
			UserName: req.UserName,
			NoKartu:  req.NoKartu,
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}
