package main

import "open-judge/src/interfaces"

type ResultDump struct {
	output  string
	err     error
	success bool
}

// GetOutput implements interfaces.Result.
func (r ResultDump) GetOutput() string {
	return r.output
}

// IsError implements interfaces.Result.
func (r ResultDump) IsError() bool {
	return !r.success
}

func (r ResultDump) GetError() error {
	return r.err
}

var _ interfaces.Result = ResultDump{}
