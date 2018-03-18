package http_p

import (
	"net/http"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/pipe"
	"github.com/21stio/go-record/pkg/types"
)

var nJobs = 100

type HttpPipe struct {
	pipe.Pipe
}

func New(p pipe.Pipe) (np HttpPipe) {
	np.Pipe = p
	return
}

func (p HttpPipe) Get(errH e.HandleError) (nP pipe.ResponsePipe) {
	nP.Ch = make(chan types.Ctx, 1000)
	nP.Scope = p.Scope

	utils.Para(p.Scope+"__http.Get", nJobs, func() {
		ctx := <-p.Ch

		rsp, err := http.Get(ctx.Val.String)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		ctx.Val.Response = rsp
		nP.Ch <- ctx
	})

	return nP
}
