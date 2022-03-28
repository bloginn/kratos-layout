package service

import (
	"context"
	"kratos-layout/internal/biz"

	pb "kratos-layout/api/open/v1"
)

type OpenService struct {
	pb.UnimplementedOpenServer
	uc *biz.OpenUseCase
}

func NewOpenService(uc *biz.OpenUseCase) *OpenService {
	return &OpenService{
		uc: uc,
	}
}

func (s *OpenService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	res, err := s.uc.Hello(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Content: res.Hello}, nil
}
