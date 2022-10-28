package top_up

import (
	"context"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/bussiness-services/app/adapter/client/top_up_services"
	"github.com/vins7/bussiness-services/app/interface/model"
	"github.com/vins7/bussiness-services/app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopUpUsecase struct {
	repo repo.TopUpRepo
}

func NewTopUpUsecase(repo repo.TopUpRepo) TopUp {
	return &TopUpUsecase{
		repo: repo,
	}
}

func (e *TopUpUsecase) TopUp(ctx context.Context, in interface{}) error {
	var req *model.TopUpReq
	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	md, err := util.GetMetadata(ctx)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	req.UserId = md.UserID
	req.UserName = md.UserName
	return e.repo.TopUp(req)
}

func (e *TopUpUsecase) Payment(ctx context.Context, in interface{}) error {
	var req *model.PaymentReq
	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	md, err := util.GetMetadata(ctx)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	req.UserId = md.UserID
	req.UserName = md.UserName
	return e.repo.Payment(req)
}
