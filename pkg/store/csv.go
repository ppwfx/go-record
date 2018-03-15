package store

import (
	"encoding/csv"
	"os"
	"log"
	"github.com/21stio/go-record/pkg/types"
)

type CsvStore struct {
	writer  *csv.Writer
	mode    os.FileMode
	columns []string
}

func Csv(path string, columns []string, mode os.FileMode) (s CsvStore) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.writer = csv.NewWriter(f)
	s.mode = mode
	s.columns = columns

	s.writer.WriteAll([][]string{
		columns,
	})

	return
}

func (s CsvStore) StoreStringMap(ctx types.Ctx, m map[string]string) (c types.Ctx, err error) {
	c = ctx

	record := []string{}
	for _, c := range s.columns {
		record = append(record, m[c])
	}

	err = s.writer.WriteAll([][]string{
		record,
	})

	return
}