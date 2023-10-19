package grpc

import (
	"context"

	"github.com/codeedu/imersao/codepix-go/application/grpc/pb"
	"github.com/codeedu/imersao/codepix-go/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) Register(ctx context.Context, in *pb.ProductRegistration) (*pb.ProductRegistrationResult, error) {
	product, err := p.ProductUseCase.RegisterProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.ProductRegistrationResult{
			Id:     "0",
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.ProductRegistrationResult{
		Id:     product.ID,
		Status: "created",
	}, nil
}

func (p *ProductGrpcService) FindById(ctx context.Context, in *pb.ProductFindById) (*pb.ProductFindByIdResult, error) {
	product, err := p.ProductUseCase.FindProductById(in.Id)

	if err != nil {
		return &pb.ProductFindByIdResult{}, err
	}

	return &pb.ProductFindByIdResult{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
