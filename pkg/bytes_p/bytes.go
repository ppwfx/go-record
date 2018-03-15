package bytes_p

import (
	"io/ioutil"
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/Masterminds/sprig"
	"html/template"
	"log"
	"github.com/21stio/go-record/pkg/strings_p"
	"github.com/21stio/go-record/pkg/store"
)

var nJobs = 10


type BytesFluid struct {
	Ch chan types.Ctx
}

func NewFluid(ctxCh chan types.Ctx) (f BytesFluid) {
	f.Ch = ctxCh

	return f
}

func (f BytesFluid) ToReadCloser(errCh chan error) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("bytes.ToReadCloser", nJobs, func() {
		ctx := <-f.Ch
		ctx.Val.ReadCloser = ioutil.NopCloser(bytes.NewReader(ctx.Val.Bytes))
		ctx.Val.Bytes = []byte{}
		out <- ctx
	})

	return out
}

func FromReadCloser(ctxCh chan types.Ctx, errCh chan error) (f BytesFluid) {
	f.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.FromReadCloser", nJobs, func() {
		ctx := <-ctxCh

		b, err := ioutil.ReadAll(ctx.Val.ReadCloser)
		if err != nil {
			errCh <- err
		}
		ctx.Val.ReadCloser.Close()

		ctx.Val.ReadCloser = nil
		ctx.Val.Bytes = b
		f.Ch <- ctx
	})

	return f
}

func Get(key string, ctxCh chan types.Ctx) (f BytesFluid) {
	f.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.Get", nJobs, func() {
		ctx := <-ctxCh

		ctx.Val.Bytes = ctx.Map.Bytes[key]

		f.Ch <- ctx
	})

	return f
}

func (f BytesFluid) HashSha512() (nF BytesFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.Hash", nJobs, func() {
		ctx := <-f.Ch

		h := sha512.New()
		h.Write(ctx.Val.Bytes)
		ctx.Val.Bytes = h.Sum(nil)

		nF.Ch <- ctx
	})

	return nF
}

func (f BytesFluid) Set(key string) (nF BytesFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.Set", nJobs, func() {
		ctx := <-f.Ch

		ctx.Map.Bytes[key] = ctx.Val.Bytes

		nF.Ch <- ctx
	})

	return
}

func (f BytesFluid) HexToString() (nF strings_p.StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.EncodeToString", nJobs, func() {
		ctx := <-f.Ch

		ctx.Val.String = hex.EncodeToString(ctx.Val.Bytes)

		ctx.Val.Bytes = []byte{}
		nF.Ch <- ctx
	})

	return
}

func (f BytesFluid) ToFile(nameTmpl string, errCh chan error) (nF BytesFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	tmpl := template.Must(template.New("a").Funcs(sprig.FuncMap()).Parse(nameTmpl))

	utils.Para("bytes.ToFile", nJobs, func() {
		ctx := <-f.Ch

		var tplOut bytes.Buffer
		tmpl.Execute(&tplOut, ctx)

		err := ioutil.WriteFile(tplOut.String(), ctx.Val.Bytes, 0644)
		if err != nil {
			errCh <- err
			return
		}

		nF.Ch <- ctx
	})

	return
}

func (f BytesFluid) Store(store store.StoreBytes, errCh chan error) (nF BytesFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.Store", nJobs, func() {
		ctx := <-f.Ch

		ctx, err := store.StoreBytes(ctx, ctx.Val.Bytes)
		if err != nil {
			errCh <- err
			return
		}

		nF.Ch <- ctx
	})

	return
}

func (f BytesFluid) IsNew(store store.IsNewBytes) (nF BytesFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("bytes.IsNew", nJobs, func() {
		ctx := <-f.Ch

		ctx, isNew, err := store.IsNewBytes(ctx, ctx.Val.Bytes)
		if err != nil {
			log.Fatal(err.Error())
		}

		if isNew {
			nF.Ch <- ctx
		}
	})

	return
}