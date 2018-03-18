package store

import (
	"github.com/21stio/go-record/pkg/types"
	"io"
)

type ValStore struct {
}

func Val() (s ValStore) {
	return
}

func (s ValStore) StoreBytes(ctx types.Ctx, b []byte) (c types.Ctx, err error) {
	c = ctx

	c.Val.Bytes = b

	return
}

func (s ValStore) StoreString(ctx types.Ctx, str string) (c types.Ctx, err error) {
	c = ctx

	c.Val.String = str

	return
}

func (s ValStore) StoreReadCloser(ctx types.Ctx, rc io.ReadCloser) (c types.Ctx, err error) {
	c = ctx

	c.Val.ReadCloser = rc

	return
}

func (s ValStore) GetReadCloser (ctx types.Ctx) (rc io.ReadCloser) {
	return ctx.Val.ReadCloser
}
