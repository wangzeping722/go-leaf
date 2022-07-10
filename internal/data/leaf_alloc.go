package data

import (
	"context"

	"go-leaf/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type leafAllocRepo struct {
	data *Data
	log  *log.Helper
}

// NewLeafAllocRepo .
func NewLeafAllocRepo(data *Data, logger log.Logger) biz.LeafAllocRepo {
	return &leafAllocRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *leafAllocRepo) Save(ctx context.Context, g *biz.LeafAlloc) (*biz.LeafAlloc, error) {
	return g, nil
}

func (r *leafAllocRepo) Update(ctx context.Context, g *biz.LeafAlloc) (*biz.LeafAlloc, error) {
	return g, nil
}

func (r *leafAllocRepo) FindByID(context.Context, int64) (*biz.LeafAlloc, error) {
	return nil, nil
}

func (r *leafAllocRepo) ListByHello(context.Context, string) ([]*biz.LeafAlloc, error) {
	return nil, nil
}

func (r *leafAllocRepo) ListAll(context.Context) ([]*biz.LeafAlloc, error) {
	return nil, nil
}
