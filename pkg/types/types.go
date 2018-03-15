package types

import (
	"io"
	"time"
)

type Ctx struct {
	CtxMeta
	CtxData
	Val CtxValue
}

type CtxValue struct {
	CtxVar
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
	ReadCloser map[string]io.ReadCloser
	Time       map[string]time.Time
	Duration   map[string]time.Duration
	Interface  map[string]interface{}
}

type CtxSlice struct {
	Strings     []string
	Ints        []int
	Bytess      [][]byte
	ReadClosers []io.ReadCloser
	Times       []time.Time
	Durations   []time.Duration
	Interfaces  []interface{}
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

type CtxVar struct {
	String     string
	Int        int
	Bytes      []byte
	ReadCloser io.ReadCloser
	Time       time.Time
	Duration   time.Duration
	Interface  interface{}
}

type Url struct {
	Url   string
	Level int
}

type Version struct {
	Hash string
	Time time.Time
}
