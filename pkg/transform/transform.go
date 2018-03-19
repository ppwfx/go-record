package transform

import (
	"log"
	"strings"
)

const (
	String           = iota
	StringSlice
)

type StringToSlice interface {
	StringToSlice (string) ([]string)
}

type ToString interface {
	ToString (string) (string)
}

type SplitTransform struct {
	Type uint
	StringSep string
}

func Split (sep interface{}) (t SplitTransform) {
	switch sep := sep.(type) {
	case string:
		t.Type = String
		t.StringSep = sep
	default:
		log.Fatal()
	}

	return
}

func (t SplitTransform) StringToSlice(str string) ([]string) {
	return strings.Split(str, t.StringSep)
}

type ReplaceTransform struct {
	Type uint
	N int
	StringOld string
	StringNew string
}

func Replace (old interface{}, new interface{}, n int) (t ReplaceTransform) {
	switch old := old.(type) {
	case string:
		new, ok := new.(string)
		if !ok {
			log.Fatal()
		}

		t.Type = String
		t.StringOld = old
		t.StringNew = new
	default:
		log.Fatal()
	}

	t.N = n

	return
}

func (t ReplaceTransform) ToString(str string) (string) {
	return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}


type AppendTransform struct {
}

func Append (suffix interface{}) (t AppendTransform) {
	return
}

func (t AppendTransform) ToString(str string) (string) {
	return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}

type PrependTransform struct {
}

func Prepend (prefix interface{}) (t PrependTransform) {
	return
}

func (t PrependTransform) ToString(str string) (string) {
	return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}

type RenderTransform struct {
}

func Render (tmpl string) (t PrependTransform) {
	return
}

func (t PrependTransform) Render(str string) (string) {
	//return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}

type IncrTransform struct {
}

func Incr (...interface{}) (t IncrTransform) {
	return
}

func (t PrependTransform) Incr(str string) (string) {
	//return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}

type BatchTransform struct {
}

func Batch (...interface{}) (t BatchTransform) {
	return
}

func (t PrependTransform) Batch(str string) (string) {
	//return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}

type SliceTransform struct {
}

func Slice (...interface{}) (t SliceTransform) {
	return
}

func (t PrependTransform) Slice(str string) (string) {
	//return strings.Replace(str, t.StringOld, t.StringNew, t.N)
}