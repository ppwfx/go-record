package utils

import (
	"github.com/21stio/go-record/pkg/types"
	"os"
)


var Debug = false

func Para(step string, n int, f func()) {
	for i := 0; i < n; i++ {
		go func() {
			for {
				f()
				if Debug {
					println(step)
				}
			}
		}()
	}
}

func Para2(step string, ctxIn chan types.Ctx, n int, passCtx bool, f func(types.Ctx, chan types.Ctx) (types.Ctx, error)) (ctxOut chan types.Ctx) {
	ctxOut = make(chan types.Ctx, 1000)

	for i := 0; i < n; i++ {
		go func() {
			for {
				ctx := <-ctxIn

				ctx, err := f(ctx, ctxOut)
				if err != nil {
					return
				}

				if passCtx {
					ctxOut <- ctx
				}

				if Debug {
					println(step)
				}
			}
		}()
	}

	return ctxOut
}

func Drain(ctxCh chan types.Ctx) {
	Para("utils.Drain", 1, func() {
		<-ctxCh
	})
}

func Do(ctxCh chan types.Ctx, f func (types.Ctx) (types.Ctx)) (ctxOut chan types.Ctx) {
	ctxOut = make(chan types.Ctx, 1000)

	Para("utils.Do", 1, func() {
		ctx := <-ctxCh

		ctx = f(ctx)

		ctxOut <- ctx
	})

	return
}

func Exit(ctxCh chan types.Ctx) {
	<-ctxCh

	os.Exit(0)
}
