package store

import (
	"github.com/21stio/go-record/pkg/types"
	"io"
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

	c.Map.Bytes[s.key] = b

	return
}

func (s CtxStore) LoadBytes(ctx types.Ctx) (c types.Ctx, err error) {
	c = ctx

	c.Val.Bytes = c.Map.Bytes[s.key]

	return
}

func (s CtxStore) StoreString(ctx types.Ctx, str string) (c types.Ctx, err error) {
	c = ctx

	c.Map.String[s.key] = str

	return
}

func (s CtxStore) LoadString(ctx types.Ctx) (c types.Ctx, err error) {
	c = ctx

	c.Val.String = c.Map.String[s.key]

	return
}

func (s CtxStore) GetString(ctx types.Ctx) (string) {
	return ctx.Map.String[s.key]
}

func (s CtxStore) StoreReadCloser(ctx types.Ctx, rc io.ReadCloser) (c types.Ctx, err error) {
	c = ctx

	c.Map.ReadCloser[s.key] = rc

	return
}

func (s CtxStore) GetReadCloser (ctx types.Ctx) (rc io.ReadCloser) {
	return ctx.Map.ReadCloser[s.key]
}
