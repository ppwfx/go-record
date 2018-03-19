package pipe

import (
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/store"
)

type ResponsePipe struct {
	Pipe
}

func (p ResponsePipe) StoreBody(s s.StoreReadCloser, errH e.HandleError) (nP ResponsePipe) {
	nP.Ch = make(chan types.Ctx, 1000)
	nP.Scope = p.Scope

	utils.Para(p.Scope+"__response.StoreBody", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := s.StoreReadCloser(ctx, ctx.Val.Response.Body)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		nP.Ch <- ctx
	})

	return nP
}
