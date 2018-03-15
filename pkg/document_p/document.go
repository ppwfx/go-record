package document_p

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/strings_p"
)

func FromReadCloser(ctxCh chan types.Ctx, errCh chan error) (out chan types.Ctx) {
	return utils.Para2("document.Get", ctxCh, 10, true, func(ctx types.Ctx, out chan types.Ctx) (types.Ctx, error) {
		doc, err := goquery.NewDocumentFromReader(ctx.Val.ReadCloser)
		if err != nil {
			errCh <- err
		}
		ctx.Val.ReadCloser.Close()

		ctx.Val.ReadCloser = nil
		ctx.Val.Interface = doc
		out <- ctx

		return ctx, err
	})
}

func ParseHrefs(ctxCh chan types.Ctx, errCh chan error) (nF strings_p.StringFluid) {
	nF.Ch = make(chan types.Ctx, 1000)

	utils.Para("string.StartsWith", 10, func() {
		ctx := <-ctxCh

		doc, ok := ctx.Val.Interface.(*goquery.Document)
		if !ok {
			return
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok {
				return
			}

			c := types.Ctx{}
			c.Parent = &ctx
			c.Id = href
			c.Val.String = href

			nF.Ch <- c
		})
	})

	return
}
