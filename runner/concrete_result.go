package runner

type concreteResult struct {
	output string
}

func (r concreteResult) String() string {
	return r.output
}

var _ Result = concreteResult{}
