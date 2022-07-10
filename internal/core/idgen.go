package core

import "context"

type IdGen interface {
	Get(ctx context.Context) (int64, error)
	Init() bool
}
