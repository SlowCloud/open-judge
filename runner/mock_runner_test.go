package runner_test

import (
	"open-judge/core"
	"open-judge/runner"
	"testing"
)

func Test_MockRunner_Run(t *testing.T) {
	runner_ := runner.MockRunner{}
	result, err := runner_.Run(runner.MockCode, core.NoLimit())

	if err != nil {
		t.Fatal("failed to compile and run code.", err)
	}
	if result.Log != runner.MockAnswer {
		t.Fatal("result is not same with expected")
	}
}

func Test_MockRunner_Run_Fail(t *testing.T) {
	runner_ := runner.MockRunner{}
	_, err := runner_.Run("random input", core.NoLimit())

	if err == nil {
		t.Fatal("should not success to run...", err)
	}
}
