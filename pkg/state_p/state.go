package state_p

//import (
//	"github.com/21stio/go-record/pkg/types"
//	"github.com/21stio/go-record/pkg/utils"
//)

var nJobs = 10

//func Set(name string, ctxIn chan types.Ctx, errCh chan error) (out chan types.Ctx) {
//	out = make(chan types.Ctx, 1000)
//
//	utils.Para("state.Set", nJobs, func() {
//		ctx := <-ctxIn
//
//		ctx.States[name] = ctx.Value
//
//		out <- ctx
//	})
//
//	return out
//}
//

//func GetBytes(name string, ctxIn chan types.Ctx, errCh chan error) (out chan types.Ctx) {
//	out = make(chan types.Ctx, 1000)
//
//	utils.Para("state.Set", nJobs, func() {
//		ctx := <-ctxIn
//
//		ctx.Value = ctx.States[name]
//
//		out <- ctx
//	})
//
//	return out
//}
