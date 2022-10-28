package top_up_client

import (
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	cfgServer "github.com/vins7/bussiness-services/config"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/top_up_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopUpClient struct {
	client proto.TopUpServicesClient
}

func NewTopUpClient(c *grpc.ClientConn) TopUpRepo {
	return &TopUpClient{
		client: proto.NewTopUpServicesClient(c),
	}
}

func (e *TopUpClient) TopUp(in interface{}) error {
	req, ok := in.(*model.TopUpReq)
	if !ok {
		return status.Errorf(codes.Internal, "Error while casting")
	}

	_, err := e.client.TopUp(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.TopUpRequest{
			UserId:   req.UserId,
			Balance:  req.Balance,
			Username: req.UserName,
			NoKartu:  req.NoKartu,
		})
	if err != nil {
		return err
	}
	return nil
}

func (e *TopUpClient) Payment(in interface{}) error {
	req, ok := in.(*model.PaymentReq)
	if !ok {
		return status.Errorf(codes.Internal, "Error while casting")
	}

	_, err := e.client.Payment(util.BuildContext(cfgServer.GetConfig().Server.App.Name),
		&proto.PaymentRequest{
			BillerId: req.BillerId,
			UserId:   req.UserId,
			Username: req.UserName,
			NoKartu:  req.NoKartu,
		})
	if err != nil {
		return err
	}
	return nil
}
