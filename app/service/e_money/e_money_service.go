package e_money_service

import (
	"context"

	uc "github.com/vins7/bussiness-services/app/usecase/e_money"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/bussiness"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EMoneyService struct {
	uc uc.EMoney
}

func NewEMoneyService(uc uc.EMoney) *EMoneyService {
	return &EMoneyService{
		uc: uc,
	}
}

func (e *EMoneyService) ListBiller(ctx context.Context, em *emptypb.Empty) (*proto.BillerResponse, error) {
	res, err := e.uc.ListBiller()
	if err != nil {
		return nil, err
	}
	return res.(*proto.BillerResponse), nil
}

func (e *EMoneyService) DetailBiller(ctx context.Context, req *proto.BillerRequest) (*proto.Biller, error) {
	res, err := e.uc.DetailBiller(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*proto.Biller), nil
}
func (e *EMoneyService) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	res, err := e.uc.GetBalance(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*proto.GetBalanceResponse), nil
}
func (e *EMoneyService) CreateEMoney(ctx context.Context, req *proto.CreateEMoneyRequest) (*proto.CreateEMoneyResponse, error) {
	res, err := e.uc.CreateEMoney(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*proto.CreateEMoneyResponse), nil
}
func (e *EMoneyService) GetTrxHist(ctx context.Context, req *proto.GetTrxHistReq) (*proto.GetTrxHistResponse, error) {
	res, err := e.uc.GetTrxHist(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*proto.GetTrxHistResponse), nil
}
