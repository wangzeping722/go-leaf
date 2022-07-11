package zero

import "context"

type ZeroGenImpl struct {
}

func NewZeroGenImpl() *ZeroGenImpl {
	return &ZeroGenImpl{}
}

func (z *ZeroGenImpl) Get(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

func (z *ZeroGenImpl) Init() bool {
	return true
}
