package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	client "github.com/vins7/bussiness-services/app/adapter/client/user_management"
	conn "github.com/vins7/bussiness-services/app/infrastructure/connection/grpc"
	cfg "github.com/vins7/bussiness-services/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer() {

	config := cfg.GetConfig()
	grpcServer := grpc.NewServer()

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
	u := client.NewUserManagementClient(conn.UserConn)
	var a interface{}
	u.Login(a)
}

func Apply(server *grpc.Server) {

}
