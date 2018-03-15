package strings_p

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"strings"
	"github.com/Masterminds/sprig"
	"html/template"
	"bytes"
	"log"
	"github.com/21stio/go-record/pkg/store"
)

var nJobs = 10

type StringFluid struct {
	Ch chan types.Ctx
}

func Pipe(ctxCh chan types.Ctx) (f StringFluid) {
	f.Ch = ctxCh

	return f
}

func (f StringFluid) Prefix(prefix string) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.Prefix", nJobs, func() {
		ctx := <-f.Ch

		ctx.Val.String = prefix + ctx.Val.String

		nF.Ch <- ctx
	})

	return
}

func (f StringFluid) StartsWith(texts []string) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.StartsWith", nJobs, func() {
		ctx := <-f.Ch

		for _, text := range texts {
			if strings.HasPrefix(ctx.Val.String, text) {
				nF.Ch <- ctx

				return
			}
		}
	})

	return
}

func (f StringFluid) Contains(texts []string) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.Contains", nJobs, func() {
		ctx := <-f.Ch

		for _, text := range texts {
			if strings.HasPrefix(ctx.Val.String, text) {
				nF.Ch <- ctx
			}
		}
	})

	return
}

func (f StringFluid) ContainsNot(text []string) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.ContainsNot", nJobs, func() {
		ctx := <-f.Ch

		for _, t := range text {
			if strings.Contains(ctx.Val.String, t) {
				return
			}
		}

		nF.Ch <- ctx
	})

	return
}

func (f StringFluid) Render(tmpl string) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	t := template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(tmpl))

	utils.Para("string.Render", nJobs, func() {
		ctx := <-f.Ch

		var tOut bytes.Buffer
		t.Execute(&tOut, ctx)

		ctx.Val.String = tOut.String()

		nF.Ch <- ctx
	})

	return
}

func (f StringFluid) Store(store store.StoreString, errCh chan error) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.Store", nJobs, func() {
		ctx := <-f.Ch

		ctx, err := store.StoreString(ctx, ctx.Val.String)
		if err != nil {
			errCh <- err
			return
		}

		nF.Ch <- ctx
	})

	return
}

func (f StringFluid) IsNew(store store.IsNewString) (nF StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.IsNew", nJobs, func() {
		ctx := <-f.Ch

		ctx, isNew, err := store.IsNewString(ctx, ctx.Val.String)
		if err != nil {
			log.Fatal(err.Error())
		}

		if isNew {
			nF.Ch <- ctx
		}
	})

	return
}
