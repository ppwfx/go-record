package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
)

type StringSlicePipe struct {
	StringSlicePipe interface{}
	Pipe
}

func (p StringSlicePipe) Each() (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string_slice.Each", nJobs, func() {
		ctx := <-p.Ch

		slice := ctx.Val.Slice.String
		ctx.Val.Slice.String = nil

		for _, str := range slice {
			ctx.Val.String = str

			np.Ch <- ctx
		}
	})

	return
}
