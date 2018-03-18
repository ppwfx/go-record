package render

import "fmt"

import (
	. "github.com/dave/jennifer/jen"
)

//"io/ioutil"
//"bytes"
//"github.com/21stio/go-record/pkg/store"
//"github.com/21stio/go-record/pkg/e"

//type BytesPipe struct {
//	Pipe
//}
//
//func (p BytesPipe) ToReadCloser() (np ReadCloserPipe) {
//	np.Ch = make(chan types.Ctx, 1000)
//	np.Scope = p.Scope
//
//	utils.Para(p.Scope+"__bytes.ToReadCloser", nJobs, func() {
//		ctx := <-p.Ch
//		ctx.Val.ReadCloser = ioutil.NopCloser(bytes.NewReader(ctx.Val.Bytes))
//		ctx.Val.Bytes = []byte{}
//		np.Ch <- ctx
//	})
//
//	return np
//}
//
//func (p BytesPipe) Load(store store.LoadBytes, errH e.HandleError) (np BytesPipe) {
//	np.Ch = make(chan types.Ctx, 1000)
//	np.Scope = p.Scope
//
//	utils.Para(p.Scope+"__bytes.Load", nJobs, func() {
//		ctx := <-p.Ch
//
//		ctx, err := store.LoadBytes(ctx)
//		if err != nil {
//			errH.HandleError(ctx, p.Ch, err)
//			return
//		}
//
//		np.Ch <- ctx
//	})
//
//	return
//}

func Run() {
	vendor := "github.com/21stio/go-record/pkg/"
	pipeType := "Bytes"
	pipeName := pipeType + "Pipe"

	f := NewFile("main")

	f.Type().Id(pipeName).Struct(
		Id("Pipe"),
	)

	f.Func().Id("main").Params().Block(
		Qual(vendor + "store", "Println").Call(Lit("Hello, world")),
	)
	f.Func().Params(
		Id("s").Id(pipeName),
	).Id("Load").Params(
		Id("store").Id("store.Store" + pipeType),
		Id("errH").Id("e.HandleError"),
	).Parens(Id("np").Id(pipeName)).Block(
		Id("np.Ch = make(chan types.Ctx, 1000)"),
		Id("np.Scope = p.Scope"),
		Return(),
	)

	fmt.Printf("%#v", f)
}




































