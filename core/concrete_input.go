package core

type concreteInput struct {
	input string
}

func (c concreteInput) String() string {
	return c.input
}

var _ Input = concreteInput{}

func NewInput(input string) Input {
	return concreteInput{input}
}

func EmptyInput() Input {
	return NewInput("")
}
