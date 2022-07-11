package service

import (
	"context"
	"strconv"

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
		log:    log.NewHelper(log.With(logger, "module", "service/leaf")),
	}
}

func (s *LeafService) Segment(ctx context.Context, req *pb.SegmentRequest) (*pb.IdReply, error) {
	id, err := s.segSvc.GetId(ctx)
	if err != nil {
		s.log.Errorf("id生成失败: %s", err.Error())
		return nil, pb.ErrorIdGenerateFailed("id生成失败")
	}
	return &pb.IdReply{Id: strconv.FormatInt(id, 10)}, nil
}
