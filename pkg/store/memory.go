package store

import (

	"github.com/21stio/go-record/pkg/maps"
	"github.com/21stio/go-record/pkg/types"
)

type MemoryStore struct {
	IsNew *maps.BoolMap
}

func Memory() (s MemoryStore) {
	s.IsNew = maps.Bool()

	return
}

func (s MemoryStore) IsNewString(ctx types.Ctx, str string) (c types.Ctx, isNew bool, err error) {
	c = ctx

	_, ok := s.IsNew.Load(str)
	if !ok {
		isNew = true
		s.IsNew.Store(str, true)
	}

	return
}

func (s MemoryStore) IsNewBytes(ctx types.Ctx, b []byte) (c types.Ctx, isNew bool, err error) {
	c = ctx

	key := string(b)

	_, ok := s.IsNew.Load(key)
	if !ok {
		isNew = true
		s.IsNew.Store(key, true)
	}

	return
}