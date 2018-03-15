package allocate_p

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"io"
	"time"
)

var nJobs = 10

func All(ctxIn chan types.Ctx) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("allocate.All", nJobs, func() {
		ctx := <-ctxIn

		ctx.Map.String     = map[string]string{}
		ctx.Map.Int        = map[string]int{}
		ctx.Map.Bytes      = map[string][]byte{}
		ctx.Map.ReadCloser = map[string]io.ReadCloser{}
		ctx.Map.Time       = map[string]time.Time{}
		ctx.Map.Duration   = map[string]time.Duration{}
		ctx.Map.Interface  = map[string]interface{}{}

		ctx.Val.Map.String     = map[string]string{}
		ctx.Val.Map.Int        = map[string]int{}
		ctx.Val.Map.Bytes      = map[string][]byte{}
		ctx.Val.Map.ReadCloser = map[string]io.ReadCloser{}
		ctx.Val.Map.Time       = map[string]time.Time{}
		ctx.Val.Map.Duration   = map[string]time.Duration{}
		ctx.Val.Map.Interface  = map[string]interface{}{}

		out <- ctx
	})

	return out
}
