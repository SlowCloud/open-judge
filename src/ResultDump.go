package main

import "open-judge/src/interfaces"

type ResultDump struct {
	output string
}

// GetOutput implements interfaces.Result.
func (r ResultDump) GetOutput() string {
	return r.output
}

var _ interfaces.Result = ResultDump{}
