package s

import (
	"encoding/csv"
	"os"
	"log"
	"github.com/21stio/go-record/pkg/types"
	"sync"
	"io"
	"bufio"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/e"
)

type CsvStore struct {
	lock    sync.RWMutex
	file    *os.File
	writer  *csv.Writer
	mode    os.FileMode
	columns []string
}

func Csv(path string, columns []string) (s CsvStore) {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.file = f
	s.columns = columns
	s.writer = csv.NewWriter(s.file)

	i, err := s.file.Stat()
	if err != nil {
		log.Fatal(err.Error())
	}

	if i.Size() != 0 {
		return
	}

	err = s.writer.WriteAll([][]string{
		s.columns,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func (s CsvStore) StoreStringMap(ctx types.Ctx, m map[string]string) (c types.Ctx, err error) {
	c = ctx

	s.lock.Lock()
	defer s.lock.Unlock()

	record := []string{}
	for _, c := range s.columns {
		record = append(record, m[c])
	}

	err = s.writer.WriteAll([][]string{
		record,
	})

	return
}

func (s CsvStore) StreamStringMap(ctxCh chan types.Ctx, errH e.HandleError) (nCtxCh chan types.Ctx) {
	nCtxCh = make(chan types.Ctx, 1000)

	nJobs := 10

	utils.Para("csv_store.StreamStringMap", nJobs, func() {
		ctx := <-ctxCh
		r := csv.NewReader(bufio.NewReader(s.file))
		s.lock.Lock()

		utils.Para("", nJobs, func() {
			r, err := r.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				s.lock.Unlock()
				errH.HandleError(ctx, ctxCh, err)
				return
			}

			m := map[string]string{}
			for i, c := range s.columns {
				m[c] = r[i]
			}

			c := ctx.NewChild()
			c.Val.Map.String = m

			nCtxCh <- c
		})
	})

	return
}
