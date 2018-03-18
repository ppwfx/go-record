package main

import (
	"log"
	"time"
	"github.com/21stio/go-record/pkg/maps"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/document_p"
	"github.com/21stio/go-record/pkg/http_p"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/allocate_p"
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/pipe"
	"encoding/hex"
	"os"
)

var urls = maps.NewConUrlMap()

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func render() (err error) {
	return
}

func run() (err error) {

	render.Run()
	os.Exit(0)

	utils.Debug = true

	urls.Store("http://www.supremenewyork.com/shop/all", types.Url{
		Url: "http://www.supremenewyork.com/shop/all",
	})

	pCh := Trigger(100 * time.Second)
	pCh = allocate_p.All(pCh)

	pipe.New(pCh).
	//Print().
		Scoped("store_html", func(p pipe.Pipe) (np pipe.Pipe) {
		np = http_p.New(p).
			Get(e.Requeue(500*time.Millisecond, true)).
			StoreBody(store.Val(), e.Fatal()).
			ReadCloser().
			ToBytes(e.Fatal()).Pipe

		np = np.Print()

		return
	}).
		Scoped("store_new_html_hash", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.Bytes().
			HashSha512().
			HexToString(hex.EncodeToString).
			IsNew(store.Memory(), e.Fatal()).
			Store(store.Ctx("hash"), e.Fatal()).Pipe

		return
	}).
		Scoped("generate_html_path", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.String().
			Render(`/tmp/pages/{{ .Id | replace "/" "_" }}_{{ index .Map.String "hash" }}.html`).
			Store(store.Ctx("path"), e.Fatal()).Pipe

		return
	}).
		Scoped("store_html", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.Bytes().
			Load(store.Ctx("html"), e.Fatal()).
			Store(store.File(store.Ctx("path"), 0644), e.Requeue(500*time.Millisecond, true)).Pipe

		return
	}).
		Scoped("store_time_url_path_in_csv", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.Do(func(ctx types.Ctx) (types.Ctx, error) {
			m := ctx.Val.Map.String
			m["time"] = time.Now().String()
			m["url"] = ctx.Id
			m["path"] = ctx.Map.String["path"]

			ctx.Val.Map.String = m

			return ctx, nil
		}, e.Fatal()).
			StringMap().
			Store(store.Csv("/tmp/pages.csv", []string{"time", "url", "path",}), e.Requeue(500*time.Millisecond, true)).Pipe

		return
	}).
		Scoped("parse_hrefs", func(p pipe.Pipe) (np pipe.Pipe) {
		p = p.Bytes().
			Load(store.Ctx("html"), e.Fatal()).
			ToReadCloser().Pipe

		np = document_p.FromReadCloser(p, store.Val(), e.Fatal()).
			ParseHrefs(e.Fatal()).Pipe

		np = np.Print()

		return
	}).
		Scoped("filter_hrefs", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.String().
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
			IsNew(store.Memory(), e.Fatal()).Pipe

		return
	}).Do(func(ctx types.Ctx) (c types.Ctx, err error) {
		urls.Store(ctx.Val.String, types.Url{
			Url:   ctx.Val.String,
			Level: ctx.Parent.Map.Int["level"] + 1})

		return
	}, e.Fatal()).
		Count("add", 10).
		Drain()

	time.Sleep(120 * time.Second)

	return
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
