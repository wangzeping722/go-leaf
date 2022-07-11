package biz

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// LeafAlloc is a LeafAlloc model.
type LeafAlloc struct {
	BizTag      string         `gorm:"column:biz_tag;primaryKey"`
	MaxId       int64          `gorm:"column:max_id"`
	Step        int64          `gorm:"column:step"`
	Description sql.NullString `gorm:"column:description"`
	UpdateTime  time.Time      `gorm:"column:update_time"`
}

type LeafAllocRepo interface {
	GetAllTags(context.Context) ([]string, error)
}

type LeafAllocUsecase struct {
	repo LeafAllocRepo
	log  *log.Helper
}

func NewLeafAllocUsecase(repo LeafAllocRepo, logger log.Logger) *LeafAllocUsecase {
	return &LeafAllocUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LeafAllocUsecase) GetAllTags(ctx context.Context) ([]string, error) {
	return uc.repo.GetAllTags(ctx)
}
