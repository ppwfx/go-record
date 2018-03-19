package pipe

import (
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/templates/store"
)

func (p __CAMEL__Pipe) Store(store store.Store__CAMEL__, errH e.HandleError) (np __CAMEL__Pipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"____SNAKE__.Store", 10, func() {
		ctx := <-p.Ch

		ctx, err := store.Store__CAMEL__(ctx, ctx.Val.String)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}