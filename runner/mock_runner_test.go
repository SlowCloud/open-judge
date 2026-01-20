package runner_test

import (
	"open-judge/runner"
	"testing"
)

func TestRun(t *testing.T) {
	mockRunner := runner.MockRunner{}

	result, err := mockRunner.Run(runner.MockInput, 1)

	if err != nil {
		t.Fatal("failed to compile and run code", err)
	}
	if result.Log != runner.MockOutput {
		t.Fatal("result is not same with expected")
	}
}

func TestRun_fail(t *testing.T) {
	mockRunner := runner.MockRunner{}

	_, err := mockRunner.Run("random input", 1)

	if err == nil {
		t.Fatal("should not success to run...", err)
	}
}
