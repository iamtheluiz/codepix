package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/iamtheluiz/codepix/codepix-go/application/grpc/pb"
	"github.com/iamtheluiz/codepix/codepix-go/application/usecase"
	"github.com/iamtheluiz/codepix/codepix-go/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	PixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, PixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
