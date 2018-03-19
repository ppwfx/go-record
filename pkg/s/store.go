package s

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
	"io"
)

type StoreBytes interface {
	StoreBytes(types.Ctx, []byte) (types.Ctx, error)
}

type StoreString interface {
	StoreString(types.Ctx, string) (types.Ctx, error)
}

type StoreStringMap interface {
	StoreStringMap(types.Ctx, map[string]string) (types.Ctx, error)
}

type StoreReadCloser interface {
	StoreReadCloser(types.Ctx, io.ReadCloser) (types.Ctx, error)
}

type GetBytes interface {
	GetBytes(types.Ctx) ([]byte)
}

type GetString interface {
	GetString(types.Ctx) (string)
}

type GetStringMap interface {
	GetStringMap(types.Ctx) (map[string]string)
}

type GetReadCloser interface {
	GetReadCloser(types.Ctx) (io.ReadCloser)
}

type LoadBytes interface {
	LoadBytes(types.Ctx) (types.Ctx, error)
}

type LoadString interface {
	LoadString(types.Ctx) (types.Ctx, error)
}

type LoadStringMap interface {
	LoadStringMap(types.Ctx) (types.Ctx, error)
}

type StreamStringMap interface {
	StreamStringMap(chan types.Ctx, e.HandleError) (chan types.Ctx)
}

type IsNewString interface {
	IsNewString(types.Ctx, string) (types.Ctx, bool, error)
}

type IsNewBytes interface {
	IsNewBytes(types.Ctx, []byte) (types.Ctx, bool, error)
}


func Parent() (i interface{}){
	return
}
