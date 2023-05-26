package service

import (
	"context"
	"encrypt-decrypt/internal/pb"
)

type LivenessProbeService struct {
	pb.UnimplementedLivenessProbeServer
}

func NewLivenessProbeService() *LivenessProbeService {
	return &LivenessProbeService{}
}

func (*LivenessProbeService) LivenessProbe(ctx context.Context, req *pb.LivenessProbeRequest) (*pb.LivenessProbeResponse, error) {
	return &pb.LivenessProbeResponse{Result: true}, nil
}
