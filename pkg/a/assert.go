package a

type String interface {
	Assert(string) bool
}

func NotEmpty () NotEmptyAssert {
	return NotEmptyAssert{}
}

type NotEmptyAssert struct {

}

func (a NotEmptyAssert) Assert(str string) (bool) {
	return str != ""
}

func IsNew (a ...interface{}) IsNewAssert {
	return IsNewAssert{}
}

type IsNewAssert struct {

}

func (a IsNewAssert) String(str string) (bool) {
	return str != ""
}



func Contains (a interface{}) ContainsAssert {
	return ContainsAssert{}
}

type ContainsAssert struct {

}

func (a ContainsAssert) String(str string) (bool) {
	return str != ""
}


func ContainsNot (a interface{}) ContainsNotAssert {
	return ContainsNotAssert{}
}

type ContainsNotAssert struct {

}

func (a ContainsNotAssert) String(str string) (bool) {
	return str != ""
}


func GreaterThan (a interface{}) ContainsNotAssert {
	return ContainsNotAssert{}
}

type GreaterThanAssert struct {

}

func (a GreaterThanAssert) String(str string) (bool) {
	return str != ""
}
