package runner

type concreteResult struct {
	output string
}

// GetOutput implements interfaces.Result.
func (r concreteResult) GetOutput() string {
	return r.output
}

var _ Result = concreteResult{}
