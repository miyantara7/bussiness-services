package e_money_client

import (
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/e_money_service"
)

type EMoneyRepo interface {
	ListBiller() (*proto.BillerResponse, error)
	DetailBiller(interface{}) (*proto.Biller, error)
	GetBalance(interface{}) (*proto.GetBalanceResponse, error)
	CreateEMoney(interface{}) (*proto.CreateEMoneyResponse, error)
	GetTrxHist(interface{}) (*proto.GetTrxHistResponse, error)
}
