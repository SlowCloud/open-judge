package runner

type concreteResult struct {
	output string
}

func (r concreteResult) GetOutput() string {
	return r.output
}

var _ Result = concreteResult{}
