package runner

type ResultDump struct {
	output string
}

// GetOutput implements interfaces.Result.
func (r ResultDump) GetOutput() string {
	return r.output
}

var _ Result = ResultDump{}
