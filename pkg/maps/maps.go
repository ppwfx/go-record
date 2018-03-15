package maps

import (
	"sync"
	"github.com/21stio/go-record/pkg/types"
)

type ConUrlMap struct {
	sync.RWMutex
	internal map[string]types.Url
}

func NewConUrlMap() *ConUrlMap {
	return &ConUrlMap{
		internal: make(map[string]types.Url),
	}
}

func (rm *ConUrlMap) Load(key string) (value types.Url, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *ConUrlMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *ConUrlMap) Store(key string, value types.Url) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func (rm *ConUrlMap) GetMap() (m map[string]types.Url) {
	rm.Lock()
	defer rm.Unlock()

	m = map[string]types.Url{}

	for k, v := range rm.internal {
		m[k] = v
	}

	return
}

func Bool() *BoolMap {
	return &BoolMap{
		internal: make(map[string]bool),
	}
}

type BoolMap struct {
	sync.RWMutex
	internal map[string]bool
}

func (rm *BoolMap) Load(key string) (value bool, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *BoolMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *BoolMap) Store(key string, value bool) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func (rm *BoolMap) GetMap() (m map[string]bool) {
	rm.Lock()
	defer rm.Unlock()

	m = map[string]bool{}

	for k, v := range rm.internal {
		m[k] = v
	}

	return
}