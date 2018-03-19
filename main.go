package main

import (
	"log"
	"time"
	"github.com/21stio/go-record/pkg/maps"
	"github.com/21stio/go-record/pkg/document_p"
	"github.com/21stio/go-record/pkg/http_p"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/e"
	"github.com/21stio/go-record/pkg/pipe"
	"encoding/hex"
	"os"
	"github.com/21stio/go-record/pkg/transform"
	"github.com/21stio/go-record/pkg/a"
	"fmt"
	"reflect"
	"github.com/21stio/go-record/pkg/s"
)

var urls = maps.NewConUrlMap()

func main() {
	err := render()
	if err != nil {
		log.Fatal(err)
	}
}

type Ent struct {
	Name     string
	Type     string
	FileName string
}



func render() (err error) {
	rootDir := "pkg/templates"
	destDir := "pkg/rendered"

	fmt.Println("%+v", reflect.ValueOf(render))

	dir := s.Ctx("dir")
	fileName := s.Ctx("fileName")
	filePath := s.Ctx("filePath")
	templateTxt := s.Ctx("templateTxt")

	ls := s.String("ls")

	utils.Debug = true

	ents := []Ent{
		{
			Name:     "ByteSlice",
			Type:     "[]byte",
			FileName: "byte_slice",
		},
		{
			Name:     "String",
			Type:     "string",
			FileName: "string",
		},
	}

	p := pipe.Drop(rootDir).
		Allocate().
		String().
		Store(dir, e.Fatal()).
		Os().
		Exec(dir, ls, e.Fatal()).
		ToSlice(transform.Split("\n")).
		Each().
		Filter(a.NotEmpty()).
		Store(fileName, e.Fatal()).
		Render(`{{ index .Map.String "dir" }}/{{ index .Map.String "fileName" }}`).Store(filePath, e.Fatal()).
		Bytes().
		Load(s.File(filePath, 0644), e.Fatal()).ToString().Store(templateTxt, e.Fatal()).Pipe

	destFilePath := s.Ctx("destFilePath")
	destFileTxt := s.Ctx("destFileTxt")

	for _, ent := range ents {
		p = p.String().
			Load(templateTxt, e.Fatal()).
			ToString(transform.Replace("type __TYPE__ interface{}", "", -1)).
			ToString(transform.Replace("__NAME__", ent.Name, -1)).
			ToString(transform.Replace("__TYPE__", ent.Type, -1)).
			Store(destFileTxt, e.Fatal()).Pipe

		p = p.String().
			Render(destDir + `/{{ index .Map.String "fileName" }}`).
			ToString(transform.Replace("__FILENAME__", ent.FileName, -1)).
			Store(destFilePath, e.Fatal()).
			Load(destFileTxt, e.Fatal()).
			ToBytes().
			Store(s.File(destFilePath, 0644), e.Fatal()).Pipe
	}

	p.Drain()

	time.Sleep(100 * time.Second)

	return
}

func run() (err error) {
	os.Exit(0)

	utils.Debug = true

	urls.Store("http://www.supremenewyork.com/shop/all", types.Url{
		Url: "http://www.supremenewyork.com/shop/all",
	})

	//pCh := Trigger(100 * time.Second)

	pipe.New().
	//Print().
		Scoped("store_html", func(p pipe.Pipe) (np pipe.Pipe) {
		np = http_p.New(p).
			Get(e.Requeue(500*time.Millisecond, true)).
			StoreBody(s.Val(), e.Fatal()).
			ReadCloser().
			ToBytes(e.Fatal()).Pipe

		np = np.Print()

		return
	}).
		Scoped("store_new_html_hash", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.Bytes().
			HashSha512().
			HexToString(hex.EncodeToString).
			IsNew(s.Memory(), e.Fatal()).
			Store(s.Ctx("hash"), e.Fatal()).Pipe

		return
	}).
		Scoped("generate_html_path", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.String().
			Render(`/tmp/pages/{{ .Id | replace "/" "_" }}_{{ index .Map.String "hash" }}.html`).
			Store(s.Ctx("path"), e.Fatal()).Pipe

		return
	}).
		Scoped("store_html", func(p pipe.Pipe) (np pipe.Pipe) {
		np = p.Bytes().
			Load(s.Ctx("html"), e.Fatal()).
			Store(s.File(s.Ctx("path"), 0644), e.Requeue(500*time.Millisecond, true)).Pipe

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
			Store(s.Csv("/tmp/pages.csv", []string{"time", "url", "path",}), e.Requeue(500*time.Millisecond, true)).Pipe

		return
	}).
		Scoped("parse_hrefs", func(p pipe.Pipe) (np pipe.Pipe) {
		p = p.Bytes().
			Load(s.Ctx("html"), e.Fatal()).
			ToReadCloser().Pipe

		np = document_p.FromReadCloser(p, s.Val(), e.Fatal()).
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
			IsNew(s.Memory(), e.Fatal()).Pipe

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
