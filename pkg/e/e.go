package e

import (
	"github.com/21stio/go-record/pkg/types"
	"log"
	"time"
)

type HandleError interface {
	HandleError(types.Ctx, chan types.Ctx, error)
}

func Fatal() FatalError {
	return FatalError{}
}

type FatalError struct {
}

func (e FatalError) HandleError(ctx types.Ctx, ctxCh chan types.Ctx, err error) {
	log.Fatal(err)
}

func Requeue(wait time.Duration, block bool) RequeueError {
	return RequeueError{
		wait:wait,
		block:block,
	}
}

type RequeueError struct {
	wait time.Duration
	block bool
}

func (e RequeueError) HandleError(ctx types.Ctx, ctxCh chan types.Ctx, err error) {
	log.Println(err)

	f := func() {
		time.Sleep(e.wait)
		ctxCh <- ctx
	}

	if e.block {
		f()
	} else {
		go f()
	}
}


