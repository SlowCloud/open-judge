package runner_test

import (
	"open-judge/core"
	"open-judge/runner"
	"testing"
)

var defaultLimit = core.Limit{TimeLimit: 1, MemoryLimit: 2000 * 1000 * 1000}

func TestRun(t *testing.T) {
	mockRunner := runner.MockRunner{}

	result, err := mockRunner.Run(runner.MockCode, defaultLimit)

	if err != nil {
		t.Fatal("failed to compile and run code", err)
	}
	if result.Log != runner.MockAnswer {
		t.Fatal("result is not same with expected")
	}
}

func TestRun_fail(t *testing.T) {
	mockRunner := runner.MockRunner{}

	_, err := mockRunner.Run("random input", defaultLimit)

	if err == nil {
		t.Fatal("should not success to run...", err)
	}
}
