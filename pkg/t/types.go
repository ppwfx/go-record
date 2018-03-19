package t

import (
	"io"
	"time"
	"net/http"
)

func (c *Ctx) NewChild() (ctx Ctx) {
	ctx.Parent = c

	return
}

type Ctx struct {
	Id         string
	Parent     *Ctx
	String     map[string]string
	Int        map[string]int
	Bytes      map[string][]byte
	Duration   map[string]time.Duration
	Interface  map[string]interface{}
	Time       map[string]time.Time
	ReadCloser map[string]io.ReadCloser
	Response   map[string]*http.Response
	Request    map[string]*http.Request
	Sliced
}

type Sliced struct {
	StringDatas     map[string][]string
	IntDatas        map[string][]int
	BytesDatas      map[string][][]byte
	ReadCloserDatas map[string][]io.ReadCloser
	TimeDatas       map[string][][]time.Time
	DurationDatas   map[string][][]time.Duration
	InterfaceDatas  map[string][]interface{}
}
