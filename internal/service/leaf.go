package service

import (
	"context"

	pb "go-leaf/api/leaf/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type LeafService struct {
	pb.UnimplementedLeafServer
	segSvc *segmentService
	log    *log.Helper
}

func NewLeafService(ss *segmentService, logger log.Logger) *LeafService {
	return &LeafService{
		segSvc: ss,
		log:    log.NewHelper(log.With(logger, "module", "service/segment")),
	}
}

func (s *LeafService) Segment(ctx context.Context, req *pb.SegmentRequest) (*pb.SegmentReply, error) {
	return &pb.SegmentReply{}, nil
}
