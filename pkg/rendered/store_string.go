package store

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
)



type StoreString interface {
	StoreString(types.Ctx, string) (types.Ctx, error)
}

type GetString interface {
	GetString(types.Ctx) (string)
}

type LoadString interface {
	LoadString(types.Ctx) (types.Ctx, error)
}

type StreamString interface {
	StreamString(chan types.Ctx, e.Handle) (chan types.Ctx)
}