package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/e"
	"os"
	"github.com/fatih/structs"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
	"io"
	"time"
)

type Pipe struct {
	Ch    chan types.Ctx
	Scope string
}

func New() (np Pipe) {
	np.Ch = make(chan types.Ctx, 1)

	go func() {
		ctx := types.Ctx{}

		np.Ch <- ctx
	}()

	return
}

func (p Pipe) Count(name string, steps int) (np Pipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope
	np.Scope = p.Scope
	c := 0

	utils.Para("inspect.Counter", nJobs, func() {
		ctx := <-p.Ch
		c++
		if c%steps == 0 {
			println(fmt.Sprintf("counter: %v %v", name, c))
		}

		np.Ch <- ctx
	})

	return
}

func Drop(any interface{}) (np Pipe) {
	np.Ch = make(chan types.Ctx, 1)

	go func() {
		ctx := types.Ctx{}
		switch sth := any.(type)  {
		case string:
			ctx.Val.String = sth
		case []string:
			ctx.Val.Slice.String = sth
		}
		np.Ch <- ctx
	}()

	return
}

func (p Pipe) Print() (np Pipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__pipe.Print", nJobs, func() {
		ctx := <-p.Ch

		//printCtx(ctx)

		spew.Dump(ctx)

		np.Ch <- ctx
	})

	return
}

func printCtx(ctx types.Ctx) {
	//ctxM := structs.Map(ctx.Val)

	//m := map[string]string{}
	//for k, v := range ctxM {
	//	fmt.Sprint(v)
	//	var s string
	//	switch v := v.(type) {
	//	case int:
	//		s = fmt.Sprint(v)
	//	case []byte:
	//		s = fmt.Sprint(string(v))
	//	case string:
	//		s = v
	//	}
	//	m[k] = s
	//}

	//m := get_map("", ctx.Val)
	//for k, v := range m {
	//	for k1, v1 := range v {
	//		if v1 == "" {
	//			delete(v, k1)
	//		}
	//	}
	//
	//	if len(m[k]) == 0 {
	//		delete(m, k)
	//	}
	//}
	//
	//for k, v := range m {
	//	keys := strings.Split(k, ":")
	//
	//	prefix := ""
	//	for _, v := range keys {
	//		println(prefix, v)
	//		prefix += "\t"
	//	}
	//
	//	for k1, v1 := range v {
	//		println(color.Black(prefix + k1 + v1))
	//	}
	//}

	//spew.Dump(ctx)
	d := strings.Split(spew.Sdump(ctx), "\n")

	d1 := []string{}
	for _, l := range d {
		if !strings.Contains(l, "<nil>") {
			d1 = append(d1, l)
		}
	}

	previousLine := ""
	containsOpen := false
	d2 := []string{}
	for _, l := range d1 {
		if containsOpen {
			if strings.Contains(l, "\t},"){
				previousLine = ""
				containsOpen = false
			} else {
				d2 = append(d2, previousLine)
				d2 = append(d2, l)
				continue
			}
		}

		if len(l) > 0 && l[len(l)-1] == '{' {
			//println(l)
			previousLine = l
			containsOpen = true
		} else {
			d2 = append(d2, l)
		}
	}

	d3 := []string{}
	for _, l := range d2 {
		if !strings.Contains(l, "Ctx") {
			d3 = append(d3, l)
		}
	}

	for _, l := range d3 {
		println(l)
	}



	//spew.Printf("%+v", ctxM)

	//return
}

func to3levels(m map[string]map[string]string) {

}

func get_map(name string, d interface{}) (m map[string]map[string]string) {
	var ctxM map[string]interface{}

	switch d := d.(type) {
	case map[string]interface{}:
		ctxM = d
		//case types.Ctx:
		//	ctxM = structs.Map(d)
		//case types.CtxValue:
		//	ctxM = structs.Map(d)
	default:
		ctxM = structs.Map(d)
	}

	m = map[string]map[string]string{}
	m[name] = map[string]string{}

	for k, v := range ctxM {
		var s string
		var cm map[string]map[string]string

		switch v := v.(type) {
		case int:
			s = fmt.Sprint(v)
		case []byte:
			s = fmt.Sprint(string(v))
		case string:
			s = v
		case map[string]interface{}:
			cname := name
			if cname != "" {
				cname += ":"
			}

			cm = get_map(cname+k, v)
		}

		//if len(cm) != 0 {
		//	spew.Dump(cm)
		//}

		for k, v := range cm {
			m[k] = v
		}

		if s != "" {
			m[name][k] = s
		}
	}

	return
}

func (p Pipe) Scoped(scope string, f func(Pipe) (Pipe)) (np Pipe) {
	if len(p.Scope) != 0 {
		scope = "__" + scope
	}
	oScope := p.Scope

	p.Scope = scope

	np = f(p)
	np.Scope = oScope

	return
}

func (p Pipe) Do(f func(types.Ctx) (types.Ctx, error), errH e.Handle) (np Pipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope+"__pipe.Do", 1, func() {
		ctx := <-p.Ch

		ctx, err := f(ctx)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		np.Ch <- ctx
	})

	return
}

func (p Pipe) Drain() {
	utils.Para(p.Scope+"__pipe.Drain", 1, func() {
		<-p.Ch
	})
}

func (p Pipe) Exit() {
	<-p.Ch

	os.Exit(0)
}

func (p Pipe) String() (np StringPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Bytes() (np BytesPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) StringMap() (np StringMapPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) ReadCloser() (np ReadCloserPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Response() (np ResponsePipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Os() (np OsPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) StringSlice() (np StringSlicePipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Allocate() (np Pipe) {
	np.Ch = make(chan types.Ctx, 1000)

	utils.Para("allocate.All", nJobs, func() {
		ctx := <-p.Ch

		ctx.Map.String     = map[string]string{}
		ctx.Map.Int        = map[string]int{}
		ctx.Map.Bytes      = map[string][]byte{}
		ctx.Map.ReadCloser = map[string]io.ReadCloser{}
		ctx.Map.Time       = map[string]time.Time{}
		ctx.Map.Duration   = map[string]time.Duration{}
		ctx.Map.Interface  = map[string]interface{}{}

		ctx.Val.Map.String     = map[string]string{}
		ctx.Val.Map.Int        = map[string]int{}
		ctx.Val.Map.Bytes      = map[string][]byte{}
		ctx.Val.Map.ReadCloser = map[string]io.ReadCloser{}
		ctx.Val.Map.Time       = map[string]time.Time{}
		ctx.Val.Map.Duration   = map[string]time.Duration{}
		ctx.Val.Map.Interface  = map[string]interface{}{}

		np.Ch <- ctx
	})

	return
}
