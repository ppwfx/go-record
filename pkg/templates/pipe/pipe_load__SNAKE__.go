package pipe

import (
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/templates/store"
)

func (p __CAMEL__Pipe) Load(store store.Load__CAMEL__, errH e.Handle) (np __CAMEL__Pipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"____SNAKE__.Load", nJobs, func() {
		ctx := <-p.Ch

		np.Ch <- ctx
	})

	return
}