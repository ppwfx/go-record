package s

import "github.com/21stio/go-record/pkg/types"

type StringStore struct {
	str string
}

func (s StringStore) GetString(ctx types.Ctx) string {
	return s.str
}

func String(str string) (StringStore) {
	return StringStore{str: str}
}