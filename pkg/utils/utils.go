package utils

import (
	"github.com/21stio/go-record/pkg/t"
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

func Para2(step string, ctxIn chan t.Ctx, n int, passCtx bool, f func(t.Ctx, chan t.Ctx) (t.Ctx, error)) (ctxOut chan t.Ctx) {
	ctxOut = make(chan t.Ctx, 1000)

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
