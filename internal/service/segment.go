package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-leaf/internal/biz"
	"go-leaf/internal/conf"
	"go-leaf/internal/core"
	"go-leaf/internal/core/segment"
	"go-leaf/internal/core/zero"
)

type segmentService struct {
	leafAllocUsecase *biz.LeafAllocUsecase
	log              *log.Helper
	idGen            core.IdGen
}

func NewSegmentService(c *conf.Leaf, leafAllocUsecase *biz.LeafAllocUsecase, logger log.Logger) *segmentService {
	ss := &segmentService{
		leafAllocUsecase: leafAllocUsecase,
		log:              log.NewHelper(log.With(logger, "module", "service/segment")),
	}

	ss.init(c, logger)
	return ss
}

func (svc *segmentService) init(c *conf.Leaf, logger log.Logger) {
	if c.Segment.Enable {
		svc.idGen = segment.NewSegmentGenImpl(svc.leafAllocUsecase, logger)
		if svc.idGen.Init() {
			svc.log.Info("Segment ID Gen Service Init Successfully")
		} else {
			svc.log.Fatal("Segment ID Gen Service Init Failed")
		}
	} else {
		svc.idGen = zero.NewZeroGenImpl()
		svc.log.Info("Zero ID Gen Service Init Successfully")
	}
}

func (svc *segmentService) GetId(ctx context.Context) (int64, error) {
	return svc.idGen.Get(ctx)
}
