package segment

import (
	"context"
	"go-leaf/internal/biz"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	"sync/atomic"
)

// 最大步数
const maxStep = 1000000

// 一个Segment维持时间为15分钟
const segmentDuration = 15 * time.Minute

var (
	ErrGeneratorInitFailed = errors.New("segment generator init failed")
	ErrKeyNotExists        = errors.New("segment key not exists")
	ErrSegmentInitFailed   = errors.New("segment init failed")
)

type SegmentGenImpl struct {
	cache            sync.Map
	leafAllocUsecase *biz.LeafAllocUsecase
	step             int
	log              *log.Helper
	initOk           int64
}

func NewSegmentGenImpl(leafAllocUsecase *biz.LeafAllocUsecase, logger log.Logger) *SegmentGenImpl {
	return &SegmentGenImpl{
		leafAllocUsecase: leafAllocUsecase,
		log:              log.NewHelper(log.With(logger, "module", "core/segment")),
	}
}

func (g *SegmentGenImpl) Get(ctx context.Context, key string) (int64, error) {
	if !g.isInitOk() {
		return 0, ErrGeneratorInitFailed
	}

	val, ok := g.cache.Load(key)
	if !ok {
		return 0, ErrKeyNotExists
	}

	sgb := val.(*Buffer)
	if !sgb.isInitOk() {
		sgb.Lock()
		if !sgb.isInitOk() {
			err := g.updateSegmentFromDb(ctx, key, sgb.getCurrentSegment())
			if err != nil {
				return 0, nil
			}

		}
		sgb.Unlock()
	}
	return 0, nil
}

func (g *SegmentGenImpl) updateSegmentFromDb(ctx context.Context, key string, seg *Segment) error {
	return nil
}

func (g *SegmentGenImpl) Init() bool {
	g.log.Info("segment generator init...")
	if err := g.updateCacheFromDb(); err != nil {
		g.log.Error(err, "segment 初始化失败")
		return false
	}

	g.setInitOk(initOk)
	g.updateCacheFromDbAtEveryMinute()
	return g.isInitOk()
}

func (g *SegmentGenImpl) updateCacheFromDbAtEveryMinute() {
	g.log.Info("start async updateCacheFromDbAtEveryMinute")
	go func() {
		timer := time.NewTimer(time.Minute)
		for {
			select {
			case <-timer.C:
				go func() {
					if err := recover(); err != nil {
						g.log.Error(err, "updateCacheFromDbAtEveryMinute panic")
					}
					err := g.updateCacheFromDb()
					if err != nil {
						g.log.Error(err, "updateCacheFromDbAtEveryMinute.updateCacheFromDb failed")
					}
				}()
			}
		}
	}()
}

func (g *SegmentGenImpl) updateCacheFromDb() error {
	st := time.Now()

	ctx := context.Background()
	dbTags, err := g.leafAllocUsecase.GetAllTags(ctx)
	if err != nil {
		return err
	}
	g.log.Info("dbTags: ", dbTags)

	if len(dbTags) == 0 {
		return nil
	}

	var cacheTags []string
	g.cache.Range(func(key, value interface{}) bool {
		cacheTags = append(cacheTags, key.(string))
		return true
	})

	insertTagsSet, removeTagsSet := make(map[string]bool), make(map[string]bool)
	for _, tag := range dbTags {
		insertTagsSet[tag] = true
	}
	for _, tag := range cacheTags {
		removeTagsSet[tag] = true
	}
	for _, tag := range cacheTags {
		if insertTagsSet[tag] {
			delete(insertTagsSet, tag)
		}
		if removeTagsSet[tag] {
			delete(removeTagsSet, tag)
		}
	}

	// 处理新增tag
	for tag := range insertTagsSet {
		buffer := newSegmentBuffer(tag)
		g.cache.Store(tag, buffer)
		g.log.Infof("Add tag %s from db to IdCache, SegmentBuffer %+v", tag, buffer)
	}

	// 处理待删除的tag
	for tag := range removeTagsSet {
		g.cache.Delete(tag)
		g.log.Infof("Del tag %s from IdCache", tag)
	}

	g.log.Infof("all tag updated, time: %s", time.Since(st))
	return nil
}

func (g *SegmentGenImpl) isInitOk() bool {
	return atomic.LoadInt64(&g.initOk) == initOk
}

func (g *SegmentGenImpl) setInitOk(i int64) {
	atomic.StoreInt64(&g.initOk, i)
}
