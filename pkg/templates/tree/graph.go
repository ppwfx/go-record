package main

import (
	"reflect"
	"github.com/21stio/go-record/pkg/e"
	"crypto/md5"
	"time"
	"github.com/21stio/go-record/pkg/t"
	"net/http"
)

const (
	Inject = iota + 1
	Eject
	Do
	Go
	Split
	Sync
	If
	Not
	Else
	Switch
)

type Signature []reflect.Type

var doSigs []Signature

type InterfaceMarker struct{}

var aInterface InterfaceMarker
var aString string
var aErrHandler e.Handle

var tInterface = reflect.TypeOf(aInterface)
var tString = reflect.TypeOf(aString)
var tErrHandler = reflect.TypeOf(aErrHandler)

func init() {
	doSigs = []Signature{
		{tInterface},
		{tInterface, tErrHandler},
		{tString, tInterface},
		{tString, tInterface, tErrHandler},
		{tString, tString, tInterface},
		{tString, tString, tInterface, tErrHandler},
	}
}

func castSignature(any []interface{}) (sig Signature) {
	for _, t := range any {
		sig = append(sig, reflect.TypeOf(t))
	}

	return
}

func isValidSignature(sig Signature, validSigs []Signature) (bool) {
	for i, s := range validSigs {
		if len(s) != len(sig) {
			continue
		}

		match := true

		for i0, t := range sig {
			if validSigs[i][i0] == reflect.TypeOf(aInterface) {
				continue
			}

			if t != validSigs[i][i0] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

type Node struct {
	Up   []*Node
	Down []*Node
	In   Signature
	Out  Signature
	Key  string
	F    interface{}
	ErrH e.Handle
}

type IPipe interface {
	Inject(any ...interface{}) (*Pipe)
	Eject(i **Pipe)
	Do(any ...interface{}) (*Pipe)
	Split(int, func([]*Pipe) (*Pipe)) (*Pipe)
	Sync(int, func([]*Pipe) (*Pipe)) (*Pipe)
	If(any ...interface{}) (*Pipe)
	Not(any ...interface{}) (*Pipe)
	Else(any ...interface{}) (*Pipe)
	Go(int) (*Pipe)
}

type Pipe struct {
	Up        []*Pipe
	Down      []*Pipe
	Kind      uint8
	Args      []interface{}
	Signature Signature
}

func (p *Pipe) Inject(any ...interface{}) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Inject
	p.Args = any
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return

	return
}

func (p *Pipe) Eject(ep *Pipe) {
	p.Kind = Eject
	p.Down = []*Pipe{ep}

	up := []*Pipe{p}
	if len(ep.Up) > 0 {
		up = append(up, ep.Up...)
	}
	ep.Up = up

	return
}

func (p *Pipe) Do(any ...interface{}) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Do
	p.Args = any
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Go(n int) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Go
	p.Args = []interface{}{n}
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Split(n int, f func([]*Pipe) (*Pipe)) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Split
	p.Args = []interface{}{n, f}
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Sync(n int, f func([]*Pipe) (*Pipe)) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Sync
	p.Args = []interface{}{n, f}
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) If(any ...interface{}) (np *Pipe) {
	np = &Pipe{}

	p.Kind = If
	p.Args = any
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Not(any ...interface{}) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Not
	p.Args = any
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Else(any ...interface{}) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Else
	p.Args = any
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func (p *Pipe) Switch(f func(interface{}) (*Pipe)) (np *Pipe) {
	np = &Pipe{}

	p.Kind = Switch
	p.Args = []interface{}{f}
	p.Down = []*Pipe{np}

	np.Up = []*Pipe{p}

	return
}

func implementsIPipe(p IPipe) {
	println("hi")
}

func traverse(p *Pipe) {
	stop := false
	for !stop {
		println(p.Kind)
		if len(p.Down) == 1 {
			p = p.Down[0]
		} else {
			stop = true
		}
	}
}

func main() {
	//p := &Pipe{}
	//b := p
	//
	//p = p.Do("url", spew.Dump).Split(1, func(pipes []*Pipe) *Pipe {
	//	return pipes[0]
	//})

	p := &Pipe{}
	b := p
	other := &Pipe{}

	loop := make(chan t.Ctx, 100)

	p.
		Inject("url", "http://www.supremenewyork.com/shop/all").
		Inject(loop).
		Do("url", http.Get).Go(100).
		Do("html").
		Do(md5.Sum).
		Split(3, func(ps []*Pipe) (p *Pipe) {
			ps[1].
				Do().
				Do("html", "").Go(100).Do()

			ps[2].
				Do("time", time.Now().String()).
				Do().
				Do("").Go(100).Do()

			return ps[0]
		}).
		Do("level").Not(true).
		Do("html", nil).Do(nil).
		If(true).
		If("").
		If("").
		Do().Do("url").
		Do("level").Do().
		Eject(other)

	traverse(b)
}
