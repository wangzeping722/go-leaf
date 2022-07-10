package biz

import (
	"context"
	"database/sql"
	"time"

	v1 "go-leaf/api/leaf/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrIdGeneratedFailed = errors.BadRequest(v1.ErrorReason_ID_GENERATE_FAILED.String(), "id generated failed")
)

// LeafAlloc is a LeafAlloc model.
type LeafAlloc struct {
	BizTag      string         `gorm:"column:biz_tag;primaryKey"`
	MaxId       int64          `gorm:"column:max_id"`
	Step        int64          `gorm:"column:step"`
	Description sql.NullString `gorm:"column:description"`
	UpdateTime  time.Time      `gorm:"column:update_time"`
}

// LeafAllocRepo is a Greater repo.
type LeafAllocRepo interface {
	Save(context.Context, *LeafAlloc) (*LeafAlloc, error)
	Update(context.Context, *LeafAlloc) (*LeafAlloc, error)
	FindByID(context.Context, int64) (*LeafAlloc, error)
	ListByHello(context.Context, string) ([]*LeafAlloc, error)
	ListAll(context.Context) ([]*LeafAlloc, error)
}

// LeafAllocUsecase is a Greeter usecase.
type LeafAllocUsecase struct {
	repo LeafAllocRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewLeafAllocUsecase(repo LeafAllocRepo, logger log.Logger) *LeafAllocUsecase {
	return &LeafAllocUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *LeafAllocUsecase) CreateGreeter(ctx context.Context, g *LeafAlloc) (*LeafAlloc, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.BizTag)
	return uc.repo.Save(ctx, g)
}
