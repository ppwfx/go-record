package document_p

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"

	"github.com/21stio/go-record/pkg/e"
	"errors"
	"github.com/21stio/go-record/pkg/pipe"
	"github.com/21stio/go-record/pkg/store"
	"log"
)

var nJobs = 100

type DocumentPipe struct {
	pipe.Pipe
}

func FromReadCloser(p pipe.Pipe, store store.GetReadCloser, errH e.HandleError) (np DocumentPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__document.FromReadCloser", nJobs, func() {
		ctx := <-p.Ch

		rc := store.GetReadCloser(ctx)
		doc, err := goquery.NewDocumentFromReader(rc)
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}
		defer rc.Close()

		ctx.Val.Interface = doc
		np.Ch <- ctx
	})

	return np
}

func (p DocumentPipe) ParseHrefs(errH e.HandleError) (np pipe.StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__document.ParseHrefs", nJobs, func() {
		ctx := <-p.Ch

		println("yo1")

		doc, ok := ctx.Val.Interface.(*goquery.Document)
		if !ok {
			errH.HandleError(ctx, p.Ch, errors.New("expected document"))
			return
		}

		println("yo2")

		h, err := doc.Html()
		if err != nil {
			log.Fatal(err)
		}

		println(h)

		println(string(ctx.Map.Bytes["html"]))

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok {
				return
			}

			c := types.Ctx{}
			c.Parent = &ctx
			c.Id = href
			c.Val.String = href

			np.Ch <- c
		})
	})

	return
}
