package pipe

import (
	"github.com/21stio/go-record/pkg/e"
	"io/ioutil"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
)

type ReadCloserPipe struct {
	Pipe
}

func (p ReadCloserPipe) ToBytes(errH e.HandleError) (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__bytes.FromReadCloser", nJobs, func() {
		ctx := <-p.Ch

		b, err := ioutil.ReadAll(ctx.Val.ReadCloser)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}
		ctx.Val.ReadCloser.Close()

		ctx.Val.ReadCloser = nil
		ctx.Val.Bytes = b
		np.Ch <- ctx
	})

	return np
}
