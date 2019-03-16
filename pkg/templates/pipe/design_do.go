package main

import (
	"reflect"
	"log"
	"github.com/21stio/go-record/pkg/t"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

const (
	Rune       = iota + 1
	Byte
	String
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Slice
	Map
	Error
	Interface
	Variadic
)

type Pipe struct {
	Meta Meta
}

type Meta struct {
	Ch        chan t.Ctx
	Up  	  []*Pipe
	Down      []*Pipe
	Allocated map[reflect.Type]bool
	In        []reflect.Type
	Out       []reflect.Type
	Key       string
	F         interface{}
	FType     reflect.Type
	DoesErr   bool
	FWrapped  interface{}
}

func (p Pipe) Inject(i *Pipe) (np Pipe) {
	return
}

func (p Pipe) Eject(i *Pipe) (np Pipe) {
	return
}

func Inject(any ...interface{}) (np Pipe) {
	np.Meta.Allocated = map[reflect.Type]bool{}

	ch, key, ctx, out := castInject(any)
	if ch != nil {
		np.Meta.Ch = ch
	} else {
		np.Meta.Key = key
		np.Meta.Out = []reflect.Type{out}
		np.Meta.Allocated[out] = true

		np.Meta.Ch = make(chan t.Ctx, 1)
		np.Meta.Ch <- ctx
	}

	return
}

func castInject(any []interface{}) (chan t.Ctx, string, t.Ctx, reflect.Type) {
	var ch chan t.Ctx
	var ctx t.Ctx
	var out reflect.Type
	var key string

	switch len(any) {
	case 0:
		log.Panic()
	case 1:
		switch sth := any[0].(type) {
		case chan t.Ctx:
			ch = sth
		default:
			key, ctx, out = CastCtx(any)
		}
	case 2:
		key, ctx, out = CastCtx(any)
	default:
		log.Panic()
	}

	return ch, key, ctx, out
}

func CastCtx(any []interface{}) (key string, ctx t.Ctx, out reflect.Type) {
	l := len(any)
	if l > 2 {
		log.Panic()
	}

	switch sth0 := any[0].(type) {
	case map[string]string:
		out = reflect.TypeOf("")
		ctx.String = map[string]string{}
		ctx.String = sth0
	case string:
		switch sth1 := any[1].(type) {
		case string:
			out = reflect.TypeOf("")
			ctx.String = map[string]string{}
			ctx.String[sth0] = sth1
			key = sth0
		default:
			log.Panic("the second type does not match")
		}
	default:
		log.Panic("the first type does not match")
	}

	if key == "" {
		log.Panic("t not set")
	}

	if out == nil {
		log.Panic("t not set")
	}

	return
}

func String_To_String(key string, f func(string) (string)) (func(t.Ctx) (t.Ctx)) {
	return func(ctx t.Ctx) (t.Ctx) {
		ctx.String[key] = f(ctx.String[key])

		return ctx
	}
}

func String_To_String_Err(key string, f func(string) (string, error)) (func(t.Ctx) (t.Ctx, error)) {
	return func(ctx t.Ctx) (t.Ctx, error) {
		var err error
		ctx.String[key], err = f(ctx.String[key])

		return ctx, err
	}
}

// todo: just execute a void function, without returning a context, maybe inject value directly in wrapper
func Interface_To_Void(key string, f func(interface{})) (func(t.Ctx) (t.Ctx)) {
	return func(ctx t.Ctx) (t.Ctx) {
		f(ctx.Interface[key])
		return ctx
	}
}

func VariadicInterface_To_Void(key string, f func(...interface{})) (func(t.Ctx) (t.Ctx)) {
	return func(ctx t.Ctx) (t.Ctx) {
		f(ctx.Interface[key])
		return ctx
	}
}

var translation map[reflect.Type]map[reflect.Type]interface{}

func init() {
	var aInterface interface{}
	var aInterfaceSlice []interface{}
	var aString string

	tString := reflect.TypeOf(aString)
	tInterface := reflect.TypeOf(aInterface)
	tInterfaceSlice := reflect.TypeOf(aInterfaceSlice)

	translation = map[reflect.Type]map[reflect.Type]interface{}{}
	translation[tString] = map[reflect.Type]interface{}{}
	translation[tString][tInterface] = func(key string, ctx t.Ctx) (t.Ctx) {
		ctx.Interface[key] = ctx.String[key]
		ctx.String[key] = ""
		return ctx
	}
	translation[tString][tInterfaceSlice] = func(key string, ctx t.Ctx) (t.Ctx) {
		// todo: proper allocation
		ctx.Interface = map[string]interface{}{}

		ctx.Interface[key] = ctx.String[key]
		ctx.String[key] = ""

		return ctx
	}
}

func translate(in reflect.Type, out reflect.Type) (interface{}) {
	return translation[in][out]
}

func (p Pipe) Do(a ...interface{}) (np Pipe) {
	np.Meta.FType, np.Meta.DoesErr, np.Meta.FWrapped = WrapFunc(p.Meta.Key, a)

	tf := translate(p.Meta.Out[0], np.Meta.FType.In(0))

	var ctx t.Ctx
	switch tf := tf.(type) {
	case func(key string, ctx t.Ctx) (t.Ctx):
		ctx = <-p.Meta.Ch
		ctx = tf(p.Meta.Key, ctx)
	case func(key string, ctx t.Ctx) (t.Ctx, error):
		var err error
		ctx = <-p.Meta.Ch
		ctx, err = tf(p.Meta.Key, ctx)
		if err != nil {
			log.Panic("ERROR case func(key string, ctx t.Ctx) (t.Ctx, error):")
		}
	default:
		log.Panic()
	}

	switch f := np.Meta.FWrapped.(type) {
	case func(t.Ctx) (t.Ctx, error):
		var err error
		ctx, err = f(ctx)
		if err != nil {
			log.Panic("ERROR case func(t.Ctx) (t.Ctx, error):")
		}
	case func(t.Ctx) (t.Ctx):
		ctx = f(ctx)
	default:
		log.Panic()
	}

	return
}

func WrapFunc(key string, af []interface{}) (fType reflect.Type, doesErr bool, fWrapped interface{}) {
	switch sf := af[0].(type) {
	case func(string) (string):
		fType = reflect.TypeOf(sf)
		fWrapped = String_To_String(key, sf)
	case func(string) (string, error):
		fType = reflect.TypeOf(sf)
		doesErr = true
		fWrapped = String_To_String_Err(key, sf)
	case func(interface{}):
		fType = reflect.TypeOf(sf)
		doesErr = false
		fWrapped = Interface_To_Void(key, sf)
	case func(...interface{}):
		fType = reflect.TypeOf(sf)
		doesErr = false
		fWrapped = VariadicInterface_To_Void(key, sf)
	default:
		log.Panic("WrapFunc did not match")
	}

	return
}

func abc(a ...interface{}) {
	println("f*** yeah!")
	spew.Dump(a)
}

func main() {
	Inject("url", "21st.io").
	Do(strings.Split)

	//injectMeta := p.Meta
	//spew.Dump(injectMeta)

	//doMeta := p.Meta
	//spew.Dump(doMeta.FType.Out(0))
}
