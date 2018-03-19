package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/e"
)

type StringMapPipe struct {
	Pipe
}

func (p StringMapPipe) Store(store s.StoreStringMap, errH e.HandleError) (np StringMapPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string_map.Store", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := store.StoreStringMap(ctx, ctx.Val.Map.String)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return np
}

func (p StringMapPipe) Stream(store s.StreamStringMap, errH e.HandleError) (np StringMapPipe) {
	np.Ch = store.StreamStringMap(p.Ch, errH)

	return
}
