package pipe

import (
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/e"
)

//func (p Pipe) String(a ...interface{}) (np Pipe) {
//	//np.Ch = p.Ch
//	//np.Scope = p.Scope
//
//	return
//}

//func (p Pipe) Bytes() (np BytesPipe) {
//	np.Ch = p.Ch
//	np.Scope = p.Scope
//
//	return
//}

func (p Pipe) StringMap() (np StringMapPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) ReadCloser() (np ReadCloserPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Response() (np ResponsePipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) Os() (np OsPipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}

func (p Pipe) StringSlice() (np StringSlicePipe) {
	np.Ch = p.Ch
	np.Scope = p.Scope

	return
}
