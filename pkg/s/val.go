package s

import (
	"github.com/21stio/go-record/pkg/t"
	"io"
)

type ValStore struct {
}

func Val() (s ValStore) {
	return
}

func (s ValStore) StoreBytes(ctx t.Ctx, b []byte) (c t.Ctx, err error) {
	c = ctx

	c.Val.Bytes = b

	return
}

func (s ValStore) StoreString(ctx t.Ctx, str string) (c t.Ctx, err error) {
	c = ctx

	c.Val.String = str

	return
}

func (s ValStore) StoreReadCloser(ctx t.Ctx, rc io.ReadCloser) (c t.Ctx, err error) {
	c = ctx

	c.Val.ReadCloser = rc

	return
}

func (s ValStore) GetReadCloser (ctx t.Ctx) (rc io.ReadCloser) {
	return ctx.Val.ReadCloser
}

func (s ValStore) GetString (ctx t.Ctx) (string) {
	return ctx.Val.String
}
