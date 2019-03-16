package s

import (
	"github.com/21stio/go-record/pkg/t"
	"io"
)

type CtxStore struct {
	key string
}

func Ctx(key string) (s CtxStore) {
	s.key = key

	return
}

func (s CtxStore) StoreBytes(ctx t.Ctx, b []byte) (c t.Ctx, err error) {
	c = ctx

	c.Bytes[s.key] = b

	return
}

func (s CtxStore) LoadBytes(ctx t.Ctx) (c t.Ctx, err error) {
	c = ctx

	//c.Val.Bytes = c.Bytes[s.key]

	return
}

func (s CtxStore) StoreString(ctx t.Ctx, str string) (c t.Ctx, err error) {
	c = ctx

	c.String[s.key] = str

	return
}

func (s CtxStore) LoadString(ctx t.Ctx) (c t.Ctx, err error) {
	c = ctx

	//c.Val.String = c.String[s.key]

	return
}

func (s CtxStore) GetString(ctx t.Ctx) (string) {
	return ctx.String[s.key]
}

func (s CtxStore) StoreReadCloser(ctx t.Ctx, rc io.ReadCloser) (c t.Ctx, err error) {
	c = ctx

	c.ReadCloser[s.key] = rc

	return
}

func (s CtxStore) GetReadCloser (ctx t.Ctx) (rc io.ReadCloser) {
	return ctx.ReadCloser[s.key]
}
