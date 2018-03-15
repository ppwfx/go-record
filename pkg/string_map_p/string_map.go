package string_map_p

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/store"
)

var nJobs = 10

type StringMapFluid struct {
	Ch chan types.Ctx
}

func Pipe(ctxCh chan types.Ctx) (f StringMapFluid) {
	f.Ch = ctxCh

	return f
}

func (f StringMapFluid)  Store(store store.StoreStringMap, errCh chan error) (nF StringMapFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string_map.Store", nJobs, func() {
		ctx := <-f.Ch

		err := store.StoreStringMap(ctx, ctx.Val.Map.String)
		if err != nil {
			errCh <- err
			return
		}

		nF.Ch <- ctx
	})

	return nF
}
