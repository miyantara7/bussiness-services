package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	client "github.com/vins7/bussiness-services/app/adapter/client/user_management_services"
	conn "github.com/vins7/bussiness-services/app/infrastructure/connection/grpc"
	svcUser "github.com/vins7/bussiness-services/app/service/user_management"
	ucUser "github.com/vins7/bussiness-services/app/usecase/user_management"
	cfg "github.com/vins7/bussiness-services/config"
	bl "github.com/vins7/module-protos/app/interface/grpc/proto/bussiness"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	clientEM "github.com/vins7/bussiness-services/app/adapter/client/emoney_services"
	svcEM "github.com/vins7/bussiness-services/app/service/e_money"
	ucEM "github.com/vins7/bussiness-services/app/usecase/e_money"

	clientTU "github.com/vins7/bussiness-services/app/adapter/client/top_up_services"
	svcTU "github.com/vins7/bussiness-services/app/service/top_up"
	ucTU "github.com/vins7/bussiness-services/app/usecase/top_up"
)

func RunServer() {

	config := cfg.GetConfig()
	grpcServer := grpc.NewServer()

	Apply(grpcServer)
	reflection.Register(grpcServer)

	svcHost := config.Server.Grpc.Host
	svcPort := config.Server.Grpc.Port

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start notpool Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}

func Apply(server *grpc.Server) {
	bl.RegisterUsermanagementServiceServer(server, svcUser.NewUserManagementService(ucUser.NewUserManagementUsecase(client.NewUserManagementClient(conn.UserConn))))
	bl.RegisterBillerServiceServer(server, svcEM.NewEMoneyService(ucEM.NewEMoneyUsecase(clientEM.NewEMoneyClient(conn.EMConn))))
	bl.RegisterTopUpServicesServer(server, svcTU.NewTopUpService(ucTU.NewTopUpUsecase(clientTU.NewTopUpClient(conn.TPConn))))
}
