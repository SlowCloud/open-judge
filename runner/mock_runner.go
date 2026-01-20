package runner

import (
	"errors"
	"open-judge/core"
)

const (
	MockInput  = "mock_input"
	MockOutput = "mock_output"
)

type MockRunner struct {
}

func (m MockRunner) Run(code core.Code, timeout int) (Result, error) {
	return m.RunWithInput(core.EmptyInput(), code, timeout)
}

func (m MockRunner) RunWithInput(input core.Input, code core.Code, timeout int) (Result, error) {
	if code.String() == MockInput {
		return Result{[]string{MockOutput}, 0, 0}, nil
	}
	return Result{}, errors.New("failed to run")
}

var _ Runner = MockRunner{}
