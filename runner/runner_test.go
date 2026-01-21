package runner_test

import (
	"open-judge/runner"
	"testing"
)

func testRun(runner_ runner.Runner, t *testing.T) {
	result, err := runner_.Run(runner.MockCode, defaultLimit)

	if err != nil {
		t.Fatal("failed to compile and run code", err)
	}
	if result.Log != runner.MockAnswer {
		t.Fatal("result is not same with expected")
	}
}

func testRun_fail(runner_ runner.Runner, t *testing.T) {
	_, err := runner_.Run("random input", defaultLimit)

	if err == nil {
		t.Fatal("should not success to run...", err)
	}
}
