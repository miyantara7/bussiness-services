package grpc

import (
	"context"
	"fmt"
	"time"

	cfg "github.com/vins7/bussiness-services/config"
	cfgServer "github.com/vins7/bussiness-services/config/server"
	"google.golang.org/grpc"
)

var (
	UserConn *grpc.ClientConn
	EMConn   *grpc.ClientConn
	TPConn   *grpc.ClientConn
)

func init() {
	var err error
	config := cfg.GetConfig()
	UserConn, err = OpenNewConnection(config.Server.UserManagement)
	if err != nil {
		fmt.Println("Not connected err =>", err.Error())
	}

	EMConn, err = OpenNewConnection(config.Server.EMoney)
	if err != nil {
		fmt.Println("Not connected err =>", err.Error())
	}

	TPConn, err = OpenNewConnection(config.Server.TopUP)
	if err != nil {
		fmt.Println("Not connected err =>", err.Error())
	}
}

func OpenNewConnection(config cfgServer.Server) (*grpc.ClientConn, error) {
	ctx, _ := context.WithTimeout(context.Background(),
		time.Duration(config.Timeout)*time.Millisecond)

	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
