package store

import (
	"github.com/21stio/go-record/pkg/types"
)

type CtxStore struct {
	key string
}

func Ctx(key string) (s CtxStore) {
	s.key = key

	return
}

func (s CtxStore) StoreBytes(ctx types.Ctx, b []byte) (c types.Ctx, err error) {
	c = ctx

	c.Map.Bytes[s.key] = c.Val.Bytes

	return
}

func (s CtxStore) StoreString(ctx types.Ctx, str string) (c types.Ctx, err error) {
	c = ctx

	c.Map.String[s.key] = c.Val.String

	return
}
