package runner_test

import (
	"open-judge/core"
	"open-judge/runner"
	"testing"
)

func Test_LocalRunner_Run(t *testing.T) {
	runner_ := runner.LocalRunner{}
	testRun(runner_, t)
}

func Test_LocalRunner_Run_Timeout(t *testing.T) {
	runner_ := runner.LocalRunner{}
	_, err := runner_.Run(`package main; import "time"; func main() {time.Sleep(10*time.Minute)}`, core.Limit{TimeLimit: 1, MemoryLimit: -1})
	if err == nil {
		t.Fatal("it must fail...")
	}
	t.Log(err)
}

func Test_LocaRunner_Run_Fail(t *testing.T) {
	runner_ := runner.LocalRunner{}
	testRun_fail(runner_, t)
}
