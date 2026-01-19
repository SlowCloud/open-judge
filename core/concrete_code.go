package core

type concreteCode struct{ code string }

func (c concreteCode) String() string {
	return c.code
}

var _ Code = concreteCode{}

func NewCode(code string) Code {
	return concreteCode{code}
}
