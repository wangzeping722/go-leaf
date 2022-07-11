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

func (l leafAllocRepo) GetAllTags(ctx context.Context) (ret []string, err error) {
	err = l.data.db.Raw("select biz_tag from leaf_alloc").Scan(&ret).Error
	if err != nil {
		l.log.Error(err)
	}
	return
}
