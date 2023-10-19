package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/codeedu/imersao/codepix-go/application/grpc/pb"
	"github.com/codeedu/imersao/codepix-go/application/usecase"
	"github.com/codeedu/imersao/codepix-go/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	productRepository := repository.ProductRepositoryDb{Db: database}
	productUsecase := usecase.ProductUseCase{ProductRepository: productRepository}
	productGrpcService := NewProductGrpcService(productUsecase)

	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start gRPC listener: ", err)
	}

	log.Printf("Starting gRPC server on %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server: ", err)
	}
}
