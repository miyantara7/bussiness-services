package e_money

import (
	"context"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/bussiness-services/app/adapter/client/emoney_services"
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/bussiness"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EMoneyUsecase struct {
	repo repo.EMoneyRepo
}

func NewEMoneyUsecase(repo repo.EMoneyRepo) EMoney {
	return &EMoneyUsecase{
		repo: repo,
	}
}

func (e *EMoneyUsecase) ListBiller() (interface{}, error) {

	res, err := e.repo.ListBiller()
	if err != nil {
		return nil, err
	}

	var out *proto.BillerResponse
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return out, nil
}

func (e *EMoneyUsecase) DetailBiller(ctx context.Context, in interface{}) (interface{}, error) {

	var req *model.BillerRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res, err := e.repo.DetailBiller(req)
	if err != nil {
		return nil, err
	}
	var out *proto.Biller
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}

func (e *EMoneyUsecase) GetBalance(ctx context.Context, in interface{}) (interface{}, error) {
	var req *model.GetBalanceRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	md, err := util.GetMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	req.UserId = md.UserID
	req.UserName = md.UserName
	res, err := e.repo.DetailBiller(req)
	if err != nil {
		return nil, err
	}

	var out *proto.GetBalanceResponse
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}

func (e *EMoneyUsecase) CreateEMoney(ctx context.Context, in interface{}) (interface{}, error) {
	var req *model.CreateEMoneyRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	md, err := util.GetMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	req.UserId = md.UserID
	req.UserName = md.UserName
	res, err := e.repo.CreateEMoney(req)
	if err != nil {
		return nil, err
	}

	var out *proto.CreateEMoneyResponse
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}

func (e *EMoneyUsecase) GetTrxHist(ctx context.Context, in interface{}) (interface{}, error) {
	var req *model.GetTrxHistReq
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	md, err := util.GetMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	req.UserId = md.UserID
	req.UserName = md.UserName
	res, err := e.repo.GetTrxHist(req)
	if err != nil {
		return nil, err
	}

	var out *proto.GetTrxHistResponse
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return out, nil
}
