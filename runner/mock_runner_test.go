package runner_test

import (
	"open-judge/core"
	"open-judge/runner"
	"testing"
)

var defaultLimit = core.Limit{TimeLimit: 1, MemoryLimit: 2000 * 1000 * 1000}

func TestRun(t *testing.T) {
	mockRunner := runner.MockRunner{}
	testRun(mockRunner, t)
}

func TestRun_fail(t *testing.T) {
	mockRunner := runner.MockRunner{}
	testRun(mockRunner, t)
}
