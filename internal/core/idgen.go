package core

type IdGen interface {
	Get() (int64, error)
}
