package types

import (
	"io"
	"time"
	"net/http"
)

type Ctx struct {
	CtxMeta
	CtxData
	Val CtxValue
}

func (c *Ctx) NewChild () (ctx Ctx) {
	ctx.Parent = c

	return
}

type CtxValue struct {
	String     string
	Int        int
	Bytes      []byte
	ReadCloser io.ReadCloser
	Response 	*http.Response
	Request 	*http.Request
	Time       time.Time
	Duration   time.Duration
	Interface  interface{}
	Map   CtxMap
	Slice CtxSlice
	NamedSlice CtxNamedSlice
}

type CtxMeta struct {
	Id     string
	Parent *Ctx
	Type   string
}

type CtxData struct {
	NamedSlice CtxSlice
	Map        CtxMap
}

type CtxMap struct {
	String     map[string]string
	Int        map[string]int
	Bytes      map[string][]byte
	Duration   map[string]time.Duration
	Interface  map[string]interface{}
	Time       map[string]time.Time
	ReadCloser map[string]io.ReadCloser
	Response 	*http.Response
	Request 	*http.Request
}

type CtxSlice struct {
	String    []string
	Int       []int
	Bytes     [][]byte
	ReadCloser[]io.ReadCloser
	Time      []time.Time
	Duration  []time.Duration
	Interface []interface{}
	Request []interface{}
	Response []interface{}
}

type CtxNamedSlice struct {
	StringDatas     map[string][]string
	IntDatas        map[string][]int
	BytesDatas      map[string][][]byte
	ReadCloserDatas map[string][]io.ReadCloser
	TimeDatas       map[string][][]time.Time
	DurationDatas   map[string][][]time.Duration
	InterfaceDatas  map[string][]interface{}
}

type Url struct {
	Url   string
	Level int
}

type Version struct {
	Hash string
	Time time.Time
}
