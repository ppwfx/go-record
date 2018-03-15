package store

import (
	"github.com/21stio/go-record/pkg/types"
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

type IsNewString interface {
	IsNewString(types.Ctx, string) (types.Ctx, bool, error)
}

type IsNewBytes interface {
	IsNewBytes(types.Ctx, []byte) (types.Ctx, bool, error)
}








