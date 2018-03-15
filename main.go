package main

import (
	"log"
	"time"
	"github.com/21stio/go-record/pkg/maps"
	"github.com/21stio/go-record/pkg/types"
	"fmt"
	"unsafe"
	"github.com/21stio/go-record/pkg/document_p"
	"github.com/21stio/go-record/pkg/http_p"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/bytes_p"
	"github.com/21stio/go-record/pkg/inspect_p"
	"github.com/21stio/go-record/pkg/allocate_p"
	"github.com/21stio/go-record/pkg/strings_p"
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/string_map_p"
)

var urls = maps.NewConUrlMap()

func main() {

	a := types.Ctx{}
	b := types.CtxValue{}
	fmt.Printf("c: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("c: %T, %d\n", b, unsafe.Sizeof(b))

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

var nJobs = 100

func run() (err error) {
	errCh := make(chan error, 100)

	//timer(time.Second)
	go printErrors(errCh)

	utils.Debug = true

	urls.Store("http://www.supremenewyork.com/shop/all", types.Url{
		Url: "http://www.supremenewyork.com/shop/all",
	})

	p := Trigger(1 * time.Second)
	p = allocate_p.All(p)
	p = http_p.Get(p, errCh)

	p = bytes_p.
		FromReadCloser(p, errCh).
		Store(store.Ctx("html"), errCh).
		HashSha512().
		IsNew(store.Memory()).
		Ch

	p = strings_p.
		Pipe(p).
		Render(`/tmp/{{ .Id | replace "/" "_" }}_{{ index .Value.StringData "version" }}.html`).
		Store(store.Ctx("path"), errCh).
		Ch

	p = bytes_p.
		Get("html", p).
		Store(store.File(`{{ index .Map.String "path" }}`, 0644), errCh).
		ToFile(`{{ index .Map.String "path" }}`, errCh).
		ToReadCloser(errCh)

	p = utils.Do(p, func(ctx types.Ctx) types.Ctx {
		m := ctx.Val.Map.String
		m["time"] = time.Now().String()
		m["url"] = ctx.Id
		m["path"] = ctx.Map.String["path"]

		ctx.Val.Map.String = m

		return ctx
	})

	p = string_map_p.
		Pipe(p).
		Store(store.Csv("/tmp/yo", []string{"time", "url", "path",}, 0644), errCh).
		Ch

	//p = inspect_p.Print(p)

	p = document_p.FromReadCloser(p, errCh)

	p = document_p.
		ParseHrefs(p, errCh).
		StartsWith([]string{
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
	}).
		ContainsNot([]string{"?"}).
		Prefix("http://www.supremenewyork.com").
		IsNew(store.Memory()).
		Ch

	p = add(p)

	p = inspect_p.Count("add", 10, p)

	//inspect_p.Print(p)

	//utils.Exit(p)

	utils.Drain(p)

	time.Sleep(120 * time.Second)

	//spew.Dump(urls)

	return
}

func printErrors(ch chan error) {
	for {
		log.Println(<-ch)
	}
}

func Trigger(frequency time.Duration) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("Trigger", 1, func() {
		for _, url := range urls.GetMap() {
			ctx := types.Ctx{}

			ctx.Id = url.Url
			ctx.Type = "go-record.url"
			ctx.Val.String = url.Url

			ctx.Map.Int = map[string]int{}
			ctx.Map.Int["level"] = url.Level

			out <- ctx
		}

		time.Sleep(frequency)
	})

	return out
}

func add(in chan types.Ctx) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("add", nJobs, func() {
		ctx := <-in

		urls.Store(ctx.Val.String, types.Url{
			Url:   ctx.Val.String,
			Level: ctx.Parent.Map.Int["level"] + 1})

		out <- ctx
	})

	return out
}
