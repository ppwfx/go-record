package http_p

import (
	"time"
	"github.com/21stio/go-record/pkg/types"
	"net/http"
	"github.com/21stio/go-record/pkg/utils"
)

var nJobs = 10

func Get(ctxIn chan types.Ctx, errCh chan error) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("http.Get", nJobs, func() {
		ctx := <-ctxIn
	R:
		rsp, err := http.Get(ctx.Val.String)
		if err != nil {
			errCh <- err
			goto R
			time.Sleep(500 * time.Millisecond)
		}

		ctx.Val.String = ""
		ctx.Val.ReadCloser = rsp.Body
		out <- ctx
	})

	return out
}
