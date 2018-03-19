package store

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
)

type __TYPE__ interface{}

type Store__CAMEL__ interface {
	Store__CAMEL__(types.Ctx, __TYPE__) (types.Ctx, error)
}

type Get__CAMEL__ interface {
	Get__CAMEL__(types.Ctx) (__TYPE__)
}

type Load__CAMEL__ interface {
	Load__CAMEL__(types.Ctx) (types.Ctx, error)
}

type Stream__CAMEL__ interface {
	Stream__CAMEL__(chan types.Ctx, e.HandleError) (chan types.Ctx)
}