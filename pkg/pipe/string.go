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
	"github.com/rs/zerolog/log"
	"github.com/21stio/go-record/pkg/transform"
	"github.com/21stio/go-record/pkg/assert"
	"github.com/21stio/go-record/pkg/s"
)

var nJobs = 10

type StringPipe struct {
	StringPipe interface{}
	Pipe
}

func (p StringPipe) Prefix(prefix string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.Prefix", nJobs, func() {
		ctx := <-p.Ch

		ctx.Val.String = prefix + ctx.Val.String

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) StartsWith(texts []string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.StartsWith", nJobs, func() {
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

func (p StringPipe) Load(store s.LoadString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.Load", nJobs, func() {
		ctx := <-p.Ch

		ctx, err := store.LoadString(ctx)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) Store(store s.StoreString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.Store", nJobs, func() {
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

func (p StringPipe) ToBytes() (np BytesPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.ToBytes", nJobs, func() {
		ctx := <-p.Ch

		ctx.Val.Bytes = []byte(ctx.Val.String)

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) ToSlice(op ...interface{}) (np StringSlicePipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	if len(op) == 0 {
		log.Fatal()
	}

	switch len(op) {
	case 0:
		log.Fatal()
	case 1:
		switch f := op[0].(type) {
		case func(string) ([]string):
			utils.Para(p.Scope+"__string.ToSlice", nJobs, func() {
				ctx := <-p.Ch

				ctx.Val.Slice.String = f(ctx.Val.String)
				ctx.Val.String = ""

				np.Ch <- ctx
			})
		case transform.StringToSlice:
			utils.Para(p.Scope+"__string.ToSlice", nJobs, func() {
				ctx := <-p.Ch

				ctx.Val.Slice.String = f.StringToSlice(ctx.Val.String)
				ctx.Val.String = ""

				np.Ch <- ctx
			})
		default:
			log.Fatal()
		}
	case 2:
		switch f := op[0].(type) {
		case func(string) ([]string, error):
			errH, ok := op[1].(e.HandleError)
			if !ok {
				log.Fatal()
				return
			}

			utils.Para(p.Scope+"__string.ToSlice", nJobs, func() {
				ctx := <-p.Ch

				slice, err := f(ctx.Val.String)
				if err != nil {
					errH.HandleError(ctx, p.Ch, err)
					return
				}

				ctx.Val.Slice.String = slice
				ctx.Val.String = ""

				np.Ch <- ctx
			})
		default:
			log.Fatal()
		}
	}

	return
}

func (p StringPipe) ToString(op ...interface{}) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	if len(op) == 0 {
		log.Fatal()
	}

	switch len(op) {
	case 0:
		log.Fatal()
	case 1:
		switch f := op[0].(type) {
		case func(string) (string):
			utils.Para(p.Scope+"__string.ToString", nJobs, func() {
				ctx := <-p.Ch

				ctx.Val.String = f(ctx.Val.String)

				np.Ch <- ctx
			})
		case transform.ToString:
			utils.Para(p.Scope+"__string.ToString", nJobs, func() {
				ctx := <-p.Ch

				ctx.Val.String = f.ToString(ctx.Val.String)

				np.Ch <- ctx
			})
		default:
			log.Fatal()
		}
	case 2:
		switch f := op[0].(type) {
		case func(string) (string, error):
			errH, ok := op[1].(e.HandleError)
			if !ok {
				log.Fatal()
				return
			}

			utils.Para(p.Scope+"__string.ToString", nJobs, func() {
				ctx := <-p.Ch

				str, err := f(ctx.Val.String)
				if err != nil {
					errH.HandleError(ctx, p.Ch, err)
					return
				}

				ctx.Val.String = str

				np.Ch <- ctx
			})
		default:
			log.Fatal()
		}
	}

	return
}

func (p StringPipe) Contains(texts []string) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.Contains", nJobs, func() {
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

	utils.Para(p.Scope+"__string.ContainsNot", nJobs, func() {
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

	utils.Para(p.Scope+"__string.Render", nJobs, func() {
		ctx := <-p.Ch

		var tOut bytes.Buffer
		t.Execute(&tOut, ctx)

		ctx.Val.String = tOut.String()

		np.Ch <- ctx
	})

	return
}

func (p StringPipe) IsNew(store s.IsNewString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.IsNew", nJobs, func() {
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

func (p StringPipe) Filter(a a.String) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__string.Filter", nJobs, func() {
		ctx := <-p.Ch

		if a.Assert(ctx.Val.String) {
			np.Ch <- ctx
		}
	})

	return
}
