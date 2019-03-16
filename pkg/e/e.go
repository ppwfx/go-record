package e

import (
	"github.com/21stio/go-record/pkg/t"
	"time"
	"log"
)

type Handle interface {
	HandleError(t.Ctx, chan t.Ctx, error)
}

func Ok() OkError {
	return OkError{}
}

type OkError struct {
}

func (e OkError) HandleError(ctx t.Ctx, ctxCh chan t.Ctx, err error) {
	log.Fatal(err)
}

func Fatal() FatalError {
	return FatalError{}
}

type FatalError struct {
}

func (e FatalError) HandleError(ctx t.Ctx, ctxCh chan t.Ctx, err error) {
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

func (e RequeueError) HandleError(ctx t.Ctx, ctxCh chan t.Ctx, err error) {
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


