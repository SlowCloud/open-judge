package runner

import (
	"errors"
)

const (
	MockInput  = "mock_input"
	MockOutput = "mock_output"
)

type MockRunner struct {
}

func (m MockRunner) Run(code string, timeout int) (Result, error) {
	return m.RunWithInput("", code, timeout)
}

func (m MockRunner) RunWithInput(input string, code string, timeout int) (Result, error) {
	if code == MockInput {
		return Result{[]string{MockOutput}, 0, 0}, nil
	}
	return Result{}, errors.New("failed to run")
}

var _ Runner = MockRunner{}
