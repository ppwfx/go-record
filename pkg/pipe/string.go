package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"strings"
	"github.com/Masterminds/sprig"
	"html/template"
	"bytes"
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/e"
)

var nJobs = 10

type StringPipe struct {
	Pipe
}

func (p StringPipe) Prefix(prefix string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.Prefix", nJobs, func() {
		ctx := <-p.Ch

		ctx.Val.String = prefix + ctx.Val.String

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) StartsWith(texts []string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.StartsWith", nJobs, func() {
		ctx := <-p.Ch

		for _, text := range texts {
			if strings.HasPrefix(ctx.Val.String, text) {
				np.Ch <- ctx

				return
			}
		}
	})

	return
}

func (p StringPipe) Contains(texts []string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.Contains", nJobs, func() {
		ctx := <-p.Ch

		for _, text := range texts {
			if strings.HasPrefix(ctx.Val.String, text) {
				np.Ch <- ctx
			}
		}
	})

	return
}

func (p StringPipe) ContainsNot(text []string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.ContainsNot", nJobs, func() {
		ctx := <-p.Ch

		for _, t := range text {
			if strings.Contains(ctx.Val.String, t) {
				return
			}
		}

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) Render(tmpl string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	t := template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(tmpl))

	utils.Para(p.Scope + "__string.Render", nJobs, func() {
		ctx := <-p.Ch

		var tOut bytes.Buffer
		t.Execute(&tOut, ctx)

		ctx.Val.String = tOut.String()

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) Store(store store.StoreString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.Store", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := store.StoreString(ctx, ctx.Val.String)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) IsNew(store store.IsNewString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.IsNew", nJobs, func() {
		ctx := <-p.Ch

		ctx, isNew, err := store.IsNewString(ctx, ctx.Val.String)
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
