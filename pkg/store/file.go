package store

import (
	"os"
	"io/ioutil"
	"github.com/21stio/go-record/pkg/types"
)

type FileStore struct {
	path GetString
	mode os.FileMode
}

func File(path GetString, mode os.FileMode) (s FileStore) {
	s.mode = mode
	s.path = path

	return
}

func (s FileStore) StoreBytes(ctx types.Ctx, b []byte) (c types.Ctx, err error) {
	c = ctx

	err = ioutil.WriteFile(s.path.GetString(ctx), ctx.Val.Bytes, 0644)
	if err != nil {
		return
	}

	return
}
