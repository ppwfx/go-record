package tree

import (
	"github.com/21stio/go-record/pkg/t"
	"reflect"
	"github.com/21stio/go-record/pkg/s"
	"time"
)

type Node struct {
	Up  	  []*Node
	Down      []*Node
	Ch        chan t.Ctx
	Allocated map[reflect.Type]bool
	In        []reflect.Type
	Out       []reflect.Type
	Key       string
	F         interface{}
	FType     reflect.Type
}


type IPipe interface {
	Inject(i *Pipe) (Pipe)
	Eject(i *Pipe)
	String(any ...interface{}) (Pipe)
	Int(any ...interface{}) (Pipe)
	Do(any ...interface{}) (Pipe)
	Split(any ...interface{}) (Pipe)
}

type Pipe struct {
	Ch        chan t.Ctx
	FWrapped  interface{}
	DoesErr   bool
}

func (p Pipe) Inject(i *Pipe) (np Pipe) {
	return
}

func (p Pipe) Eject(i *Pipe) (np Pipe) {
	return
}

func Inject(any ...interface{}) (np Pipe) {

}

np.Meta.Allocated = map[reflect.Type]bool{}

ch, key, ctx, out := castInject(any)
if ch != nil {
np.Meta.Ch = ch
} else {
np.Meta.Key = key
np.Meta.Out = []reflect.Type{out}
np.Meta.Allocated[out] = true

np.Meta.Ch = make(chan t.Ctx, 1)
np.Meta.Ch <- ctx
}

return


p.
	Inject(map[string]string{
	"url": "http://www.supremenewyork.com/shop/all",
	}).
	Inject(loop).
	String("url", http.Get, 100).
	Bytes("html").
	Do(md5.Sum).
	If(a.IsNew(s.Memory())).
	String("path", transform.Render(`/tmp/pages/{{ index .String "url" | replace "/" "_" }}_{{ index .String "hash" }}.html`)).
	Split(3, func(ps []Pipe) (p Pipe) {
		ps[1].
			Do(Buffer(100)).
			Bytes("html", s.File(s.Ctx("path"), 0644)).Go(100).Do(Drain)

		ps[2].
			String("time", time.Now().String()).
			Do(Buffer(100)).
			Do(s.Csv("/tmp/urls.csv", []string{"time", "path", "url"})).Go(100).Do(Drain)

		return ps[0]
	}).
	Int("level").Not(a.GreaterThan(3)).
	Bytes("html", parseHrefs).Do(transform.Slice()).
	If(a.IsNew(s.Memory())).
	If(a.Contains([]string{
	"/shop/jackets",
	"/shop/shirts",
	"/shop/tops_sweaters",
	"/shop/sweatshirts",
	"/shop/pants",
	"/shop/shorts",
	"/shop/hats",
	"/shop/bags",
	"/shop/accessories",
	"/shop/shoes",
	"/shop/skate",
	})).
	If(a.ContainsNot("?")).
	Do(transform.Prepend("http://www.supremenewyork.com")).String("url").
	Int("level", s.Parent()).Do(transform.Incr).
	Eject(loop)