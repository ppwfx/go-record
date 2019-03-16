package s

import "github.com/21stio/go-record/pkg/t"

type StringStore struct {
	str string
}

func (s StringStore) GetString(ctx t.Ctx) string {
	return s.str
}

func String(str string) (StringStore) {
	return StringStore{str: str}
}