package util

import (
	"context"

	"github.com/vins7/bussiness-services/app/interface/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func BuildContext(client string) (ctx context.Context) {
	return metadata.AppendToOutgoingContext(context.Background(), "client-name", client)
}

func GetMetadata(ctx context.Context) (model.MetaData, error) {
	var metaData model.MetaData
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return model.MetaData{}, status.Errorf(codes.InvalidArgument, "UnaryEcho: failed to get metadata")
	}

	if t, ok := md["userid"]; ok {
		metaData.UserID = t[0]
	}

	if t, ok := md["username"]; ok {
		metaData.UserName = t[0]
	}

	return metaData, nil
}
