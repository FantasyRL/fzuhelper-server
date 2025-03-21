/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type AtomicStrings struct {
	v atomic.Value
}

func NewAtomicStrings() *AtomicStrings {
	as := &AtomicStrings{}
	as.v.Store([]string{})
	return as
}

func (as *AtomicStrings) Load() []string {
	//nolint:forcetypeassert
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
