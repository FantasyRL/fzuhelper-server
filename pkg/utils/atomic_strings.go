package utils

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type AtomicStrings struct {
	v atomic.Value
}

// 初始化时先存一个空切片
func newAtomicStrings() *AtomicStrings {
	as := &AtomicStrings{}
	as.v.Store([]string{})
	return as
}

func (as *AtomicStrings) Load() []string {
	return as.v.Load().([]string)
}

func (as *AtomicStrings) Store(val []string) {
	as.v.Store(val)
}

func (as *AtomicStrings) Len() int {
	slice := as.Load()
	return len(slice)
}

func (as *AtomicStrings) Random() (string, bool) {
	slice := as.Load()
	n := len(slice)
	if n == 0 {
		return "", false
	}

	idx := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
	return slice[idx], true
}
