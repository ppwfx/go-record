package s

import (
	"github.com/21stio/go-record/pkg/t"
	"github.com/21stio/go-record/pkg/e"
	"io"
)

type StoreBytes interface {
	StoreBytes(t.Ctx, []byte) (t.Ctx, error)
}

type StoreString interface {
	StoreString(t.Ctx, string) (t.Ctx, error)
}

type StoreStringMap interface {
	StoreStringMap(t.Ctx, map[string]string) (t.Ctx, error)
}

type StoreReadCloser interface {
	StoreReadCloser(t.Ctx, io.ReadCloser) (t.Ctx, error)
}

type GetBytes interface {
	GetBytes(t.Ctx) ([]byte)
}

type GetString interface {
	GetString(t.Ctx) (string)
}

type GetStringMap interface {
	GetStringMap(t.Ctx) (map[string]string)
}

type GetReadCloser interface {
	GetReadCloser(t.Ctx) (io.ReadCloser)
}

type LoadBytes interface {
	LoadBytes(t.Ctx) (t.Ctx, error)
}

type LoadString interface {
	LoadString(t.Ctx) (t.Ctx, error)
}

type LoadStringMap interface {
	LoadStringMap(t.Ctx) (t.Ctx, error)
}

type StreamStringMap interface {
	StreamStringMap(chan t.Ctx, e.Handle) (chan t.Ctx)
}

type IsNewString interface {
	IsNewString(t.Ctx, string) (t.Ctx, bool, error)
}

type IsNewBytes interface {
	IsNewBytes(t.Ctx, []byte) (t.Ctx, bool, error)
}


func Parent() (i interface{}){
	return
}
