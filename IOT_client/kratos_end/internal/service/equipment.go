package service

import (
	"context"
	"kratos/internal/biz"

	pb "kratos/api/equipment/v1"
)

type EquipmentService struct {
	pb.UnimplementedEquipmentServer
	uc *biz.WlProductsRepo
}

func NewEquipmentService(uc *biz.WlProductsRepo) *EquipmentService {
	return &EquipmentService{
		uc: uc,
	}
}

func (s *EquipmentService) ProductsList(ctx context.Context, req *pb.ProductsListReq) (*pb.ProductsListResp, error) {
	return &pb.ProductsListResp{}, nil
}
