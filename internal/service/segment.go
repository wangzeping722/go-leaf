package service

import (
	"context"
	"go-leaf/internal/biz"
	"go-leaf/internal/conf"
	"go-leaf/internal/core"
	"go-leaf/internal/core/zero"

	"github.com/go-kratos/kratos/v2/log"
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

	ss.init(c)
	return ss
}

func (svc *segmentService) init(c *conf.Leaf) {
	svc.idGen = zero.NewZeroGenImpl()
}

func (svc *segmentService) GetId(ctx context.Context) (int64, error) {
	return svc.idGen.Get(ctx)
}
