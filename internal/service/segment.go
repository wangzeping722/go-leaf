package service

import (
	"go-leaf/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type segmentService struct {
	leafAllocUsecase *biz.LeafAllocUsecase
	log              *log.Helper
}

func NewSegmentService(leafAllocUsecase *biz.LeafAllocUsecase, logger log.Logger) *segmentService {
	ss := &segmentService{
		leafAllocUsecase: leafAllocUsecase,
		log:              log.NewHelper(log.With(logger, "module", "service/segment")),
	}

	ss.init()
	return ss
}

func (ssvc *segmentService) init() {

}
