package store

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
)



type StoreByteSlice interface {
	StoreByteSlice(types.Ctx, []byte) (types.Ctx, error)
}

type GetByteSlice interface {
	GetByteSlice(types.Ctx) ([]byte)
}

type LoadByteSlice interface {
	LoadByteSlice(types.Ctx) (types.Ctx, error)
}

type StreamByteSlice interface {
	StreamByteSlice(chan types.Ctx, e.Handle) (chan types.Ctx)
}