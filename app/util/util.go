package util

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func BuildContext(client string) (ctx context.Context) {
	return metadata.AppendToOutgoingContext(context.Background(), "client-name", client)
}
