package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
)

type IStringPipe interface {
	Print() (IStringPipe)
	Inject(string) (IStringPipe)
	Do(...interface{}) (IStringPipe)
	If(...interface{}) (IStringPipe)
}

//string
//string error
//ctx
//ctx err
//
//string transformer