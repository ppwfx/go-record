package logic_p

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
)

var nJobs = 10

func Split(ctxCh chan types.Ctx) (out0 chan types.Ctx, out1 chan types.Ctx) {
	out0 = make(chan types.Ctx, 1000)
	out1 = make(chan types.Ctx, 1000)

	utils.Para("logic.Split", nJobs, func() {
		ctx := <-ctxCh
		out0 <- ctx
		out1 <- ctx
	})

	return
}