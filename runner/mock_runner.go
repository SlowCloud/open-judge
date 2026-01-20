package runner

import (
	"errors"
	"open-judge/core"
)

const (
	MockCode   = "mock_code"
	MockAnswer = "mock_answer"
)

type MockRunner struct {
}

func (m MockRunner) Run(code string, limit core.Limit) (Result, error) {
	return m.RunWithInput(code, "", limit)
}

func (m MockRunner) RunWithInput(code string, input string, limit core.Limit) (Result, error) {
	if code == MockCode {
		return Result{MockAnswer, 0, 0}, nil
	}
	return Result{}, errors.New("failed to run")
}

var _ Runner = MockRunner{}
