package pipe

import (
	"io/ioutil"
	"crypto/sha512"
	"encoding/hex"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/e"
	"bytes"
)

type BytesPipe struct {
	Pipe
}

func (p BytesPipe) ToReadCloser() (np ReadCloserPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.ToReadCloser", nJobs, func() {
		ctx := <-p.Ch
		ctx.Val.ReadCloser = ioutil.NopCloser(bytes.NewReader(ctx.Val.Bytes))
		ctx.Val.Bytes = []byte{}
		np.Ch <- ctx
	})

	return np
}

func (p BytesPipe) Load(store store.LoadBytes, errH e.HandleError) (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.Load", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := store.LoadBytes(ctx)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}

func (p BytesPipe) HashSha512() (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.HashSha512", nJobs, func() {
		ctx := <-p.Ch

		h := sha512.New()
		h.Write(ctx.Val.Bytes)
		ctx.Val.Bytes = h.Sum(nil)

		np.Ch <- ctx
	})

	return np
}

func (p BytesPipe) HexToString(func([]byte) string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.EncodeToString", nJobs, func() {
		ctx := <-p.Ch

		ctx.Val.String = hex.EncodeToString(ctx.Val.Bytes)

		ctx.Val.Bytes = []byte{}
		np.Ch <- ctx
	})

	return
}

func (p BytesPipe) Store(store store.StoreBytes, errH e.HandleError) (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.Store", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := store.StoreBytes(ctx, ctx.Val.Bytes)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}

func (p BytesPipe) IsNew(store store.IsNewBytes, errH e.HandleError) (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__bytes.IsNew", nJobs, func() {
		ctx := <-p.Ch

		ctx, isNew, err := store.IsNewBytes(ctx, ctx.Val.Bytes)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		if isNew {
			np.Ch <- ctx
		}
	})

	return
}
