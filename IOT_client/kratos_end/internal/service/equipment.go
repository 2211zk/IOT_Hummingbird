package service

import (
	"context"

	pb "kratos/api/equipment/v1"
)

type EquipmentService struct {
	pb.UnimplementedEquipmentServer
}

func NewEquipmentService() *EquipmentService {
	return &EquipmentService{}
}

func (s *EquipmentService) ProductsList(ctx context.Context, req *pb.ProductsListReq) (*pb.ProductsListResp, error) {
	return &pb.ProductsListResp{}, nil
}
