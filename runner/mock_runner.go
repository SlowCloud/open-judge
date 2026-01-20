package runner

import (
	"errors"
)

const (
	MockCode   = "mock_code"
	MockAnswer = "mock_answer"
)

type MockRunner struct {
}

func (m MockRunner) Run(code string, timeout int) (Result, error) {
	return m.RunWithInput("", code, timeout)
}

func (m MockRunner) RunWithInput(input string, code string, timeout int) (Result, error) {
	if code == MockCode {
		return Result{MockAnswer, 0, 0}, nil
	}
	return Result{}, errors.New("failed to run")
}

var _ Runner = MockRunner{}
