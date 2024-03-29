package pipe

import (
	"github.com/21stio/go-record/pkg/e"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"errors"
	"time"
	"crypto/sha512"
	"github.com/21stio/go-record/pkg/a"
	"github.com/21stio/go-record/pkg/transform"
	"os"
	"reflect"
	"github.com/21stio/go-record/pkg/t"
	"crypto/md5"
	"github.com/21stio/go-record/pkg/s"
	"go/ast"
)
type Pipe struct {
	Up []*Pipe
	Down []*Pipe
	Kind uint8
	FType reflect.Type
	DoesErr bool
	F func(t.Ctx)(t.Ctx)
	FErr func(t.Ctx)(t.Ctx, error)
}

type IPipe interface {
	Print() (IPipe)
	Named() (IPipe)
	In(any interface{}) (IPipe)
	Do(func(t.Ctx) (t.Ctx, error), e.Handle) (IPipe)
	Drain() (IPipe)
	Exit() (IPipe)
}

func (p Pipe) String(a ...interface{}) (np Pipe) {
	return
}

func (p Pipe) Bytes(a ...interface{}) (np Pipe) {
	return
}






func (p Pipe) Out(chan t.Ctx) (np Pipe) {
	return
}
func (p Pipe) In(a ...interface{}) (np Pipe) {
	return
}

// 1st(*bool, assert.) 2nd( func(Pipe)(Pipe))
func (p Pipe) If(a ...interface{}) (np Pipe) {
	return
}
func (p Pipe) Do(a ...interface{}) (np Pipe) {
	return
}

func (p Pipe) Go(n ...interface{}) (np Pipe) {
	return
}

func (p Pipe) As(n ...string) (np Pipe) {
	return
}

func (p Pipe) Not(a ...interface{}) (np Pipe) {
	return
}

// 1st([]any) 2nd( func(map[any]Pipe)(Pipe))
func (p Pipe) Switch(a ...interface{}) (np Pipe) {
	return
}




func (p Pipe) Split(int, func([]Pipe)(Pipe)) (np Pipe) {
	return
}

func (p Pipe) Each(a ...interface{}) (np Pipe) {
	return
}


func parseHrefs (ctx t.Ctx, doc *goquery.Document) (out chan t.Ctx, err error) {
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if !ok {
			err = errors.New("attr does not contain href")

			return
		}

		c := ctx.NewChild()

		c.String["href"] = href

		out <- c
	})

	return
}

func parseHrefs2 (doc *goquery.Document) (hrefs []string, err error) {
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if !ok {
			err = errors.New("attr does not contain href")

			return
		}

		hrefs = append(hrefs, href)
	})

	return//
}


func yoyo ([]byte) []byte {
	sha512.New()
}

func Buffer(int) (interface{}) {
	os.Getenv("IS_MASTER")
	return
}

func Drain() (interface{}) {
	return
}


func abc() {
	p := Pipe{}

	loop := make(chan t.Ctx, 100)

	p.
		In(map[string]string{
			"url": "http://www.supremenewyork.com/shop/all",
		}).
		In(loop).
		Do("url", http.Get).Go(100).
		As("html").
		Do(md5.Sum).
		If(a.IsNew(s.Memory())).
		Do(transform.Render(`/tmp/pages/{{ index .String "url" | replace "/" "_" }}_{{ index .String "hash" }}.html`), "path").
		Split(3, func(ps []Pipe) (p Pipe) {
			ps[1].
				Do(Buffer(10000)).
				Do("html", s.File(s.Ctx("path"), 0644)).Go(100).Do(Drain)

			ps[2].
				Do(time.Now, "time").
				Do(Buffer(100)).
				Do(s.Csv("/tmp/urls.csv", []string{"time", "path", "url"})).Go(100).Do(Drain)

			return ps[0]
		}).
		Not("level", a.GreaterThan(3)).
		Do("html", parseHrefs).
		Each(func(g Pipe) {
			g.
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
				Do("level", s.Parent()).Do(transform.Incr).
				Out(loop)
		})
}