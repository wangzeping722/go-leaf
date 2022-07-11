package segment

import (
	"sync"
	"sync/atomic"
)

const initOk = 1
const nextReady = 1

type Segment struct {
	value  int64
	max    int64
	step   int
	buffer *Buffer
}

func newSegment(buf *Buffer) *Segment {
	return &Segment{
		buffer: buf,
	}
}

func (s *Segment) getNewValue() int64 {
	return atomic.AddInt64(&s.value, 1)
}

func (s *Segment) getCurValue() int64 {
	return atomic.LoadInt64(&s.value)
}

func (s *Segment) setVal(v int64) {
	atomic.StoreInt64(&s.value, v)
}

func (s *Segment) getMax() int64 {
	return atomic.LoadInt64(&s.max)
}

func (s *Segment) getIdle() int64 {
	return atomic.LoadInt64(&s.max)
}

type Buffer struct {
	sync.RWMutex
	key        string
	segments   []*Segment
	currentPos int64
	nextReady  int64
	initOk     int64

	step            int64
	minStep         int64
	updateTimeStamp int64
}

func newSegmentBuffer(key string) *Buffer {
	buffer := &Buffer{
		key: key,
	}
	buffer.segments = []*Segment{newSegment(buffer), newSegment(buffer)}
	return buffer
}

func (b *Buffer) isInitOk() bool {
	return atomic.LoadInt64(&b.initOk) == initOk
}

func (b *Buffer) setInitOk() {
	atomic.StoreInt64(&b.initOk, initOk)
}

func (b *Buffer) isNextReady() bool {
	return atomic.LoadInt64(&b.nextReady) == nextReady
}

func (b *Buffer) setNextReady() {
	atomic.StoreInt64(&b.nextReady, nextReady)
}

func (b *Buffer) getCurrentSegment() *Segment {
	return b.segments[b.getCurrentPos()]
}

func (b *Buffer) getCurrentPos() int64 {
	return atomic.LoadInt64(&b.currentPos)
}

func (b *Buffer) getNextPos() int64 {
	return (atomic.LoadInt64(&b.currentPos) + 1) % 2
}

func (b *Buffer) switchPos() {
	atomic.StoreInt64(&b.currentPos, b.getNextPos())
}

func (b *Buffer) getUpdateTimestamp() int64 {
	return atomic.LoadInt64(&b.updateTimeStamp)
}

func (b *Buffer) setUpdateTimestamp(ts int64) {
	atomic.StoreInt64(&b.updateTimeStamp, ts)
}
